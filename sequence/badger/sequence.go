package badger

import (
	"log"

	"github.com/dgraph-io/badger"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/sequence"
)

type SequenceBadger struct {
	badgerDB       *badger.DB
	badgerSequence *badger.Sequence
}

func (badgerSeq *SequenceBadger) Open() (err error) {
	log.Printf("Connecting sequence write to %v at %v\n", conf.BADGER, conf.Conf.SequenceBadger.Dir)
	opts := badger.DefaultOptions
	opts.Dir = conf.Conf.SequenceBadger.Dir
	opts.ValueDir = conf.Conf.SequenceBadger.ValueDir
	opts.SyncWrites = conf.Conf.SequenceBadger.SyncWrites
	badgerSeq.badgerDB, err = badger.Open(opts)
	if err != nil {
		log.Printf("Sequence badger open error: %v", err)
		return err
	}
	keyname := []byte(conf.Conf.SequenceBadger.KeyName)

	badgerSeq.badgerSequence, err = badgerSeq.badgerDB.GetSequence(keyname, conf.Conf.SequenceBadger.Bandwidth)
	if err != nil {
		log.Printf("Sequence badger sequence error: %v", err)
		badgerSeq.badgerDB.Close()
		return err
	}
	return nil
}

func (badgerSeq *SequenceBadger) Close() {
	if badgerSeq.badgerSequence != nil {
		badgerSeq.badgerSequence.Release()
		badgerSeq.badgerSequence = nil
	}
	if badgerSeq.badgerDB != nil {
		badgerSeq.badgerDB.Close()
		badgerSeq.badgerDB = nil
	}
}

func (badgerSeq *SequenceBadger) NextSequence() (lastID uint64, err error) {
	lastID, err = badgerSeq.badgerSequence.Next()
	if err != nil {
		log.Printf("Sequence badger get error: %v", err)
		return 0, err
	}
	log.Printf("Sequence badger LastID=%v", lastID)
	return lastID, nil
}

var badgerSeq = SequenceBadger{}

func init() {
	sequence.MustRegister(string(conf.BADGER), &badgerSeq)
}
