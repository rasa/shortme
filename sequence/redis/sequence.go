package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/rasa/shortme/sequence"
)

type SequenceRedis struct {
	redisClient *redis.Client
}

func (redisSeq *SequenceRedis) Open() (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Printf("Sequence redis open error: %v", err)
		return err
	}

	redisSeq.redisClient = client
	return nil
}

func (redisSeq *SequenceRedis) Close() {
	if redisSeq.redisClient != nil {
		redisSeq.redisClient.Close()
		redisSeq.redisClient = nil
	}
}

func (redisSeq *SequenceRedis) NextSequence() (sequence uint64, err error) {
	keyName := "shortme"
	err = redisSeq.redisClient.Incr(keyName).Err()
	if err != nil {
		log.Printf("Sequence redis incr error: %v", err)
		return 0, err
	}
	var lastID int64
	lastID, err = redisSeq.redisClient.Get(keyName).Int64()
	if err != nil {
		log.Printf("Sequence redis get error: %v", err)
		return 0, err
	}

	log.Printf("Sequence redis LastID=%v", lastID)

	sequence = uint64(lastID)
	// mysql sequence will start at 1, we actually want it to be
	// started at 0. :)
	sequence -= 1
	return sequence, nil
}

var redisSeq = SequenceRedis{}

func init() {
	log.Printf("Register sequence %v", "db")
	sequence.MustRegister("redis", &redisSeq)
}
