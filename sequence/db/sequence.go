package db

import (
	"database/sql"
	"log"

	"github.com/andyxning/shortme/conf"
	"github.com/andyxning/shortme/sequence"
	_ "github.com/go-sql-driver/mysql"
)

type SequenceDB struct {
	db *sql.DB
}

func (dbSeq *SequenceDB) Open() (err error) {
	var db *sql.DB
	db, err = sql.Open("mysql", conf.Conf.SequenceDB.DSN)
	if err != nil {
		log.Printf("sequence db open error. %v", err)
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
	return nil
}

func (dbSeq *SequenceDB) Close() {
	if dbSeq.db != nil {
		dbSeq.db.Close()
		dbSeq.db = nil
	}
}

func (dbSeq *SequenceDB) NextSequence() (sequence uint64, err error) {
	var stmt *sql.Stmt
	stmt, err = dbSeq.db.Prepare(`UPDATE sequence SET id=LAST_INSERT_ID(id+1) LIMIT 1`)
	if err != nil {
		log.Printf("sequence db prepare error. %v", err)
		return 0, err
	}
	defer stmt.Close()

	var res sql.Result
	res, err = stmt.Exec()
	if err != nil {
		log.Printf("sequence db replace into error. %v", err)
		return 0, err
	}

	// 兼容LastInsertId方法的返回值
	var lastID int64
	lastID, err = res.LastInsertId()
	if err != nil {
		log.Printf("sequence db get LastInsertId error. %v", err)
		return 0, err
	} else {
		sequence = uint64(lastID)
		// mysql sequence will start at 1, we actually want it to be
		// started at 0. :)
		sequence -= 1
		return sequence, nil
	}
}

var dbSeq = SequenceDB{}

func init() {
	sequence.MustRegister("db", &dbSeq)
}
