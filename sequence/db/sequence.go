package db

import (
	"database/sql"
	"log"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/sequence"
)

const sequenceSQL = "UPDATE sequence SET id=LAST_INSERT_ID(id+1)"

// these would work, too:
// const sequenceSQL = "INSERT INTO sequence (id) VALUES (NULL)"
// const sequenceSQL = "REPLACE INTO  sequence (id) VALUES (NULL)"

type SequenceDB struct {
	db   *sql.DB
	stmt *sql.Stmt
}

func (dbSeq *SequenceDB) Open() (err error) {
	re := regexp.MustCompile("@([^/]*)")
	b := re.FindStringSubmatch(conf.Conf.SequenceDB.DSN)
	log.Printf("Connecting sequence write to %v at %v\n", conf.MYSQL, b[0])
	db, err := sql.Open("mysql", conf.Conf.SequenceDB.DSN)
	if err != nil {
		log.Printf("Sequence db open error: %v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("sequence db ping error. %v", err)
		return err
	}

	db.SetMaxIdleConns(conf.Conf.SequenceDB.MaxIdleConns)
	db.SetMaxOpenConns(conf.Conf.SequenceDB.MaxOpenConns)

	dbSeq.db = db

	dbSeq.stmt, err = dbSeq.db.Prepare(sequenceSQL)
	if err != nil {
		log.Printf("Sequence db prepare error: %v", err)
		return err
	}

	return nil
}

func (dbSeq *SequenceDB) Close() {
	if dbSeq.stmt != nil {
		dbSeq.stmt.Close()
		dbSeq.stmt = nil
	}
	if dbSeq.db != nil {
		dbSeq.db.Close()
		dbSeq.db = nil
	}
}

func (dbSeq *SequenceDB) NextSequence() (sequence uint64, err error) {
	res, err := dbSeq.stmt.Exec()
	if err != nil {
		log.Printf("Sequence db update error: %v", err)
		return 0, err
	}

	// 兼容LastInsertId方法的返回值
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Printf("Sequence db LastInsertId error: %v", err)
		return 0, err
	}

	// mysql sequence will start at 1, we actually want it to be
	// started at 0. :)
	sequence = uint64(lastID) - 1
	return sequence, err
}

var dbSeq = SequenceDB{}

func init() {
	sequence.MustRegister(string(conf.MYSQL), &dbSeq)
}
