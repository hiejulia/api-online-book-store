package clients

import (
	"github.com/avast/retry-go"
	"github.com/hiejulia/api-online-book-store/utils"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Constants
var (
	RedisDB         int
	RedisHost       string
	RedisPort       string
	RedisPass       string
	configRedisOnce sync.Once
	redisOnce       sync.Once
	redisConnOnce   sync.Once
)

// Redis ...
type Redis struct {
	cfg  Config
	conn *redis.Client
}

// RedisConfig ...
func RedisConfig() Config {
	configRedisOnce.Do(func() {
		RedisDB = utils.GetEnvInt("CACHE_DB")
		RedisHost = utils.GetEnvStr("CACHE_HOST")
		RedisPort = utils.GetEnvStr("CACHE_PORT")
		RedisPass = utils.GetEnvStr("CACHE_PASS")

		if RedisDB < 0 {
			RedisDB = 0
		}
	})

	return Config{
		DB:       RedisDB,
		Host:     RedisHost,
		Password: RedisPass,
		Port:     RedisPort,
	}
}

// NewRedis ...
func NewRedis(cfg Config) (red *Redis) {
	redisOnce.Do(func() {
		red = &Redis{cfg: cfg}
	})
	return
}

// Close ...
func (red *Redis) Close() (err error) {
	err = red.conn.Close()
	return
}

// Get ...
func (red *Redis) Get(key string) (val string, err error) {
	err = retry.Do(
		func() error {
			var redisErr error
			val, redisErr = red.conn.Get(key).Result()
			return redisErr
		},
		retry.Attempts(3),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("[Get] Retrying request after error: %v", err)
		}),
	)
	if err != nil {
		log.Printf("[Get] Redis get req err: %v for key %s", err, key)
	}
	return
}

func (red *Redis) GetInt(key string) (val int64, err error) {
	err = retry.Do(
		func() error {
			var redisErr error
			val, redisErr = red.conn.Get(key).Int64()
			return redisErr
		},
		retry.Attempts(3),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("[GetInt] Retrying request after error: %v", err)
		}),
	)
	if err != nil {
		log.Printf("[GetInt] Redis GetInt req err: %v for key %s", err, key)
	}
	return
}

// Open ...
func (red *Redis) Open() (err error) {
	redisConnOnce.Do(func() {
		opt := &redis.Options{
			Addr:     red.cfg.Addr(),
			Password: red.cfg.Password,
			DB:       red.cfg.DB,
		}
		red.conn = redis.NewClient(opt)
	})
	return
}

// Set ...
func (red *Redis) Set(key, val string, exp time.Duration) error {
	var err error
	err = retry.Do(
		func() error {
			redisErr := red.conn.Set(key, val, exp).Err()
			return redisErr
		},
		retry.Attempts(3),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("[Set] Retrying request after error: %v", err)
		}),
	)
	if err != nil {
		log.Printf("[Set] Redis set req err: %v for key %s", err, key)
	}
	return err
}

func (red *Redis) SetInt(key string, val int64, exp time.Duration) error {
	var err error
	err = retry.Do(
		func() error {
			redisErr := red.conn.Set(key, val, exp).Err()
			return redisErr
		},
		retry.Attempts(3),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("[SetInt] Retrying request after error: %v", err)
		}),
	)
	if err != nil {
		log.Printf("[SetInt] Redis SetInt req err: %v for key %s", err, key)
	}
	return err
}
