package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/sequence"
)

type SequenceRedis struct {
	redisClient *redis.Client
}

func (redisSeq *SequenceRedis) Open() (err error) {
	log.Printf("Connecting sequence write to %v at %v\n", conf.REDIS, conf.Conf.SequenceRedis.Addr)
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Conf.SequenceRedis.Addr,
		Password:     conf.Conf.SequenceRedis.Password,
		PoolSize:     conf.Conf.SequenceRedis.PoolSize,
		DialTimeout:  10 * time.Second, // * time.Duration(conf.Conf.SequenceRedis.DialTimeout),
		PoolTimeout:  30 * time.Second, // * time.Duration(conf.Conf.SequenceRedis.PoolTimeout),
		ReadTimeout:  30 * time.Second, // * time.Duration(conf.Conf.SequenceRedis.ReadTimeout),
		WriteTimeout: 30 * time.Second, // * time.Duration(conf.Conf.SequenceRedis.WriteTimeout),
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
	err = redisSeq.redisClient.Incr(conf.Conf.SequenceRedis.KeyName).Err()
	if err != nil {
		log.Printf("Sequence redis incr error: %v", err)
		return 0, err
	}
	sequence, err = redisSeq.redisClient.Get(conf.Conf.SequenceRedis.KeyName).Uint64()
	if err != nil {
		log.Printf("Sequence redis get error: %v", err)
		return 0, err
	}

	log.Printf("Sequence redis sequence=%v", sequence)
	return sequence, err
}

var redisSeq = SequenceRedis{}

func init() {
	sequence.MustRegister(string(conf.REDIS), &redisSeq)
}
