package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/sequence"
)

type SequenceDB struct {
	db   *sql.DB
	stmt *sql.Stmt
}

func (dbSeq *SequenceDB) Open() (err error) {
	var db *sql.DB
	db, err = sql.Open("mysql", conf.Conf.SequenceDB.DSN)
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

	dbSeq.stmt, err = dbSeq.db.Prepare(`UPDATE sequence SET id=LAST_INSERT_ID(id+1)`)
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
	var res sql.Result
	res, err = dbSeq.stmt.Exec()
	if err != nil {
		log.Printf("Sequence db update error: %v", err)
		return 0, err
	}

	// 兼容LastInsertId方法的返回值
	var lastID int64
	lastID, err = res.LastInsertId()
	if err != nil {
		log.Printf("Sequence db LastInsertId error: %v", err)
		return 0, err
	}

	sequence = uint64(lastID)
	// mysql sequence will start at 1, we actually want it to be
	// started at 0. :)
	sequence -= 1
	return sequence, nil
}

var dbSeq = SequenceDB{}

func init() {
	log.Printf("Register sequence %v", "db")
	sequence.MustRegister("db", &dbSeq)
}
