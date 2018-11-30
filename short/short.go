package short

import (
	"database/sql"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/dchest/siphash"
	"github.com/go-sql-driver/mysql"
	"github.com/minio/highwayhash"
	"github.com/rasa/shortme/base"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/sequence"

	_ "github.com/rasa/shortme/sequence/badger"
	_ "github.com/rasa/shortme/sequence/db"
	_ "github.com/rasa/shortme/sequence/redis"
)

type HashType int

const (
	NoHash HashType = iota
	HighwayHash
	SipHash
)

const (
	hashType               = HighwayHash
	highwayhash_key_string = "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
	siphash_k0             = uint64(316665572293978160)
	siphash_k1             = uint64(8573005253291875333)
	reconnect_tries        = 2
)

type shorter struct {
	readDB          *sql.DB
	writeDB         *sql.DB
	sequence        sequence.Sequence
	selectLongStmt  *sql.Stmt
	selectShortStmt *sql.Stmt
	insertStmt      *sql.Stmt
}

var (
	ErrQueryShortDB      = errors.New("Short read db query error")
	ErrScanShortRows     = errors.New("Short read db query rows scan error")
	ErrIterShortRows     = errors.New("Short read db query rows iterate error")
	ErrGetNextSeq        = errors.New("Get next sequence error")
	ErrPrepareSQL        = errors.New("Short write db prepares error")
	ErrInsertData        = errors.New("Short write db insert error")
	ErrBadScheme         = errors.New("Bad scheme")
	ErrBadHost           = errors.New("Bad host")
	ErrFailedToDecodeKey = errors.New("Failed to decode key")
)

var highwayhash_key []byte

func init() {
	var err error
	highwayhash_key, err = hex.DecodeString(highwayhash_key_string)
	if err != nil {
		log.Panicf("Failed to decode key: %v", err)
	}
}

// connect will panic when it can not connect to DB server.
func (shorter *shorter) mustConnect() {
	shorter.reconnectReadDB()
	shorter.reconnectWriteDB()
}

func (shorter *shorter) readClose() {
	if shorter.selectLongStmt != nil {
		shorter.selectLongStmt.Close()
		shorter.selectLongStmt = nil
	}
	if shorter.selectShortStmt != nil {
		shorter.selectShortStmt.Close()
		shorter.selectShortStmt = nil
	}
	if shorter.readDB != nil {
		shorter.readDB.Close()
		shorter.readDB = nil
	}
}

func (shorter *shorter) reconnectReadDB() {
	shorter.readClose()

	re := regexp.MustCompile("@([^/]*)")
	b := re.FindStringSubmatch(conf.Conf.ShortDB.ReadDSN)
	log.Printf("Connecting short read to %v at %v\n", conf.MYSQL, b[0])
	db, err := sql.Open("mysql", conf.Conf.ShortDB.ReadDSN)
	if err != nil {
		log.Panicf("Short read db open error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicf("Short read db ping error: %v", err)
	}

	db.SetMaxIdleConns(conf.Conf.ShortDB.MaxIdleConns)
	db.SetMaxOpenConns(conf.Conf.ShortDB.MaxOpenConns)

	shorter.readDB = db

	selectLongSQL := "SELECT long_url FROM short WHERE short_url=?"
	shorter.selectLongStmt, err = shorter.readDB.Prepare(selectLongSQL)
	if err != nil {
		log.Panicf("Short db prepare long error: %v", err)
	}

	selectShortSQL := "SELECT short_url FROM short WHERE long_hash=? and long_url=?"
	shorter.selectShortStmt, err = shorter.readDB.Prepare(selectShortSQL)
	if err != nil {
		log.Panicf("Short db prepare short error: %v", err)
	}
}

func (shorter *shorter) writeClose() {
	if shorter.insertStmt != nil {
		shorter.insertStmt.Close()
		shorter.insertStmt = nil
	}
	if shorter.writeDB != nil {
		shorter.writeDB.Close()
		shorter.writeDB = nil
	}
}

func (shorter *shorter) reconnectWriteDB() {
	shorter.writeClose()
	re := regexp.MustCompile("@([^/]*)")
	b := re.FindStringSubmatch(conf.Conf.ShortDB.WriteDSN)
	log.Printf("Connecting short write to %v at %v\n", conf.MYSQL, b[0])
	db, err := sql.Open("mysql", conf.Conf.ShortDB.WriteDSN)
	if err != nil {
		log.Panicf("Short write db open error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicf("Short write db ping error: %v", err)
	}

	db.SetMaxIdleConns(conf.Conf.ShortDB.MaxIdleConns)
	db.SetMaxOpenConns(conf.Conf.ShortDB.MaxOpenConns)

	shorter.writeDB = db

	insertSQL := "INSERT INTO short(long_url, short_url, long_hash) VALUES(?, ?, ?)"

	shorter.insertStmt, err = shorter.writeDB.Prepare(insertSQL)
	if err != nil {
		log.Panicf("Short write db prepare error: %v", err)
	}
}

// initSequence will panic when it can not open the sequence successfully.
func (shorter *shorter) mustInitSequence() {
	shorter.reconnectSequence()
}

func (shorter *shorter) sequenceClose() {
	if shorter.sequence != nil {
		shorter.sequence.Close()
		shorter.sequence = nil
	}
}

func (shorter *shorter) reconnectSequence() {
	shorter.sequenceClose()

	log.Printf("Connecting to sequence write via %v\n", conf.Conf.SequenceBackend)
	seq, err := sequence.GetSequence(string(conf.Conf.SequenceBackend))
	if err != nil {
		log.Panicf("Invalid backend type: %v: %v", conf.Conf.SequenceBackend, err)
	}

	err = seq.Open()
	if err != nil {
		log.Panicf("Open sequence instance error: %v", err)
	}

	shorter.sequence = seq
}

func (shorter *shorter) close() {
	shorter.readClose()
	shorter.writeClose()
	shorter.sequenceClose()
}

func (shorter *shorter) connectionOK(err *error) bool {
	if *err == nil {
		return true
	}
	if *err == driver.ErrBadConn {
		return false
	}
	if *err == mysql.ErrInvalidConn {
		return false
	}
	return true
}

func (shorter *shorter) Expand(shortURL string) (longURL string, err error) {
	retry := 0
	for retry < reconnect_tries {
		err = shorter.selectLongStmt.QueryRow(shortURL).Scan(&longURL)
		if shorter.connectionOK(&err) {
			break
		}
		shorter.reconnectReadDB()
		retry += 1
	}
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No longURL found for %v", shortURL)
		return "", nil
	case err != nil:
		log.Printf("Short read db query error: %v", err)
		return "", ErrQueryShortDB
	}
	return longURL, nil
}

func (shorter *shorter) Short(longURL string) (shortURL string, err error) {
	var long_hash uint64
	var u *url.URL

	longURL = strings.TrimSpace(longURL)
	u, err = url.Parse(longURL)
	if err != nil {
		log.Printf("URL parse error %v: %v", err, longURL)
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	u.Scheme = strings.ToLower(u.Scheme)
	if u.Scheme != "http" && u.Scheme != "https" {
		log.Printf("URL bad scheme: %v", longURL)
		return "", ErrBadScheme
	}
	if u.Host == "" {
		log.Printf("URL bad host: %v", longURL)
		return "", ErrBadHost
	}
	u.Host = strings.ToLower(u.Host)
	longURL = u.String()

	switch hashType {
	case HighwayHash:
		hash, err := highwayhash.New64(highwayhash_key)
		if err != nil {
			log.Printf("Failed to decode key: %v", err)
			return "", ErrFailedToDecodeKey
		}
		hash.Write([]byte(longURL))
		long_hash = hash.Sum64()
	case SipHash:
		long_hash = siphash.Hash(siphash_k0, siphash_k1, []byte(longURL))
	case NoHash:
		long_hash = 0
	}

	var retry int = 0

	if hashType != NoHash {
		short_url := ""
		for retry < reconnect_tries {
			err = shorter.selectShortStmt.QueryRow(long_hash, longURL).Scan(&short_url)
			if shorter.connectionOK(&err) {
				break
			}
			shorter.reconnectReadDB()
			retry += 1
		}

		switch {
		case err == sql.ErrNoRows:
			log.Printf("No shortURL found for %v (%v)", long_hash, longURL)
		case err != nil:
			log.Printf("short read db query error. %v", err)
			return "", ErrQueryShortDB
		default:
			if short_url != "" {
				return short_url, nil
			}
		}
	}

	for {
		var seq uint64

		retry = 0
		for retry < reconnect_tries {
			seq, err = shorter.sequence.NextSequence()
			if shorter.connectionOK(&err) {
				break
			}
			shorter.reconnectSequence()
			retry += 1
		}

		if err != nil {
			log.Printf("get next sequence error. %v", err)
			return "", ErrGetNextSeq
		}

		shortURL = base.Int2String(seq)
		if _, exists := conf.Conf.Common.BlackShortURLsMap[shortURL]; !exists {
			break
		}
	}

	retry = 0
	for retry < reconnect_tries {
		_, err = shorter.insertStmt.Exec(longURL, shortURL, long_hash)
		if shorter.connectionOK(&err) {
			break
		}
		shorter.reconnectWriteDB()
		retry += 1
	}

	if err != nil {
		log.Printf("short write db insert error. %v", err)
		return "", ErrInsertData
	}

	return shortURL, nil
}

var Shorter shorter

func Start() {
	Shorter.mustConnect()
	Shorter.mustInitSequence()
	log.Println("shorter starts")
}

func Close() {
	log.Println("shorter closes")
	Shorter.close()
}
