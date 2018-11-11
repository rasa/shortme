package short

import (
	"database/sql"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/dchest/siphash"
	"github.com/go-sql-driver/mysql"
	"github.com/minio/highwayhash"
	"github.com/rasa/shortme/base"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/sequence"
	_ "github.com/rasa/shortme/sequence/db"
)

const (
	UseHighwayHash         = true
	highwayhash_key_string = "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
	siphash_k0             = uint64(316665572293978160)
	siphash_k1             = uint64(8573005253291875333)
)

type shorter struct {
	readDB   *sql.DB
	writeDB  *sql.DB
	sequence sequence.Sequence
}

var (
	ErrQueryShortDB  = errors.New("short read db query error")
	ErrScanShortRows = errors.New("short read db query rows scan error")
	ErrIterShortRows = errors.New("short read db query rows iterate error")
	ErrGetNextSeq    = errors.New("get next sequence error")
	ErrPrepareSQL    = errors.New("short write db prepares error")
	ErrInsertData    = errors.New("short write db insert error")
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

func (shorter *shorter) reconnectReadDB() {
	if shorter.readDB != nil {
		shorter.readDB.Close()
		shorter.readDB = nil
	}

	db, err := sql.Open("mysql", conf.Conf.ShortDB.ReadDSN)
	if err != nil {
		log.Panicf("short read db open error. %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicf("short read db ping error. %v", err)
	}

	db.SetMaxIdleConns(conf.Conf.ShortDB.MaxIdleConns)
	db.SetMaxOpenConns(conf.Conf.ShortDB.MaxOpenConns)

	shorter.readDB = db
}

func (shorter *shorter) reconnectWriteDB() {
	if shorter.writeDB != nil {
		shorter.writeDB.Close()
		shorter.writeDB = nil
	}

	db, err := sql.Open("mysql", conf.Conf.ShortDB.WriteDSN)
	if err != nil {
		log.Panicf("short write db open error. %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicf("short write db ping error. %v", err)
	}

	db.SetMaxIdleConns(conf.Conf.ShortDB.MaxIdleConns)
	db.SetMaxOpenConns(conf.Conf.ShortDB.MaxOpenConns)

	shorter.writeDB = db
}

// initSequence will panic when it can not open the sequence successfully.
func (shorter *shorter) mustInitSequence() {
	shorter.reconnectSequence()
}

func (shorter *shorter) reconnectSequence() {
	if shorter.sequence != nil {
		shorter.sequence.Close()
		shorter.sequence = nil
	}

	seq, err := sequence.GetSequence("db")
	if err != nil {
		log.Panicf("get sequence instance error. %v", err)
	}

	err = seq.Open()
	if err != nil {
		log.Panicf("open sequence instance error. %v", err)
	}

	shorter.sequence = seq
}

func (shorter *shorter) close() {
	if shorter.readDB != nil {
		shorter.readDB.Close()
		shorter.readDB = nil
	}

	if shorter.writeDB != nil {
		shorter.writeDB.Close()
		shorter.writeDB = nil
	}
}

func (shorter *shorter) connectionError(err *error) bool {
	return *err == driver.ErrBadConn || *err == mysql.ErrInvalidConn
}

func (shorter *shorter) Expand(shortURL string) (longURL string, err error) {
	selectSQL := fmt.Sprintf(`SELECT long_url FROM short WHERE short_url=?`)

	var rows *sql.Rows
	rows, err = shorter.readDB.Query(selectSQL, shortURL)

	if err != nil {
		if shorter.connectionError(&err) {
			shorter.reconnectReadDB()

			rows, err = shorter.readDB.Query(selectSQL, shortURL)
			if err != nil {
				log.Printf("short read db query error. %v", err)
				return "", ErrQueryShortDB
			}
		} else {
			log.Printf("short read db query error. %v", err)
			return "", ErrQueryShortDB
		}
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&longURL)
		if err != nil {
			log.Printf("short read db query rows scan error. %v", err)
			return "", ErrScanShortRows
		}
	}

	err = rows.Err()
	if err != nil {
		log.Printf("short read db query rows iterate error. %v", err)
		return "", ErrIterShortRows
	}

	return longURL, nil
}

func (shorter *shorter) Short(longURL string) (shortURL string, err error) {
	var long_hash uint64

	if UseHighwayHash {
		hash, err := highwayhash.New64(highwayhash_key)
		if err != nil {
			log.Printf("Failed to decode key: %v", err)
			return "", errors.New("Failed to decode key")
		}
		hash.Write([]byte(longURL))
		long_hash = hash.Sum64()
	} else {
		long_hash = siphash.Hash(siphash_k0, siphash_k1, []byte(longURL))
	}
	selectSQL := fmt.Sprintf(`SELECT short_url FROM short WHERE long_hash=? and long_url=?`)

	var rows *sql.Rows
	rows, err = shorter.readDB.Query(selectSQL, long_hash, longURL)
	if err != nil {
		log.Printf("short read db query error. %v", err)
		return "", errors.New("short read db query error")
	}

	defer rows.Close()

	short_url := ""
	for rows.Next() {
		err = rows.Scan(&short_url)
		if err != nil {
			log.Printf("short read db query rows scan error. %v", err)
			return "", errors.New("short read db query rows scan error")
		}
	}

	err = rows.Err()
	if err != nil {
		log.Printf("short read db query rows iterate error. %v", err)
		return "", errors.New("short read db query rows iterate error")
	}

	if short_url != "" {
		return short_url, nil
	}

	for {
		var seq uint64
		seq, err = shorter.sequence.NextSequence()
		if err != nil {
			if shorter.connectionError(&err) {
				shorter.reconnectSequence()

				seq, err = shorter.sequence.NextSequence()
				if err != nil {
					log.Printf("get next sequence error. %v", err)
					return "", ErrGetNextSeq
				}
			} else {
				log.Printf("get next sequence error. %v", err)
				return "", ErrGetNextSeq
			}
		}

		shortURL = base.Int2String(seq)
		if _, exists := conf.Conf.Common.BlackShortURLsMap[shortURL]; !exists {
			break
		}
	}

	insertSQL := fmt.Sprintf(`INSERT INTO short(long_url, short_url, long_hash) VALUES(?, ?, ?)`)

	var stmt *sql.Stmt
	stmt, err = shorter.writeDB.Prepare(insertSQL)
	if err != nil {
		if shorter.connectionError(&err) {
			shorter.reconnectWriteDB()

			stmt, err = shorter.writeDB.Prepare(insertSQL)
			if err != nil {
				log.Printf("short write db prepares error. %v", err)
				return "", ErrPrepareSQL
			}
		} else {
			log.Printf("short write db prepares error. %v", err)
			return "", ErrPrepareSQL
		}
	}
	defer stmt.Close()

	_, err = stmt.Exec(longURL, shortURL, long_hash)
	if err != nil {
		if shorter.connectionError(&err) {
			shorter.reconnectWriteDB()

			_, err = stmt.Exec(longURL, shortURL)
			if err != nil {
				log.Printf("short write db insert error. %v", err)
				return "", ErrInsertData
			}
		} else {
			log.Printf("short write db insert error. %v", err)
			return "", ErrInsertData
		}
	}

	return shortURL, nil
}

var Shorter shorter

func Start() {
	Shorter.mustConnect()
	Shorter.mustInitSequence()
	log.Println("shorter starts")
}
