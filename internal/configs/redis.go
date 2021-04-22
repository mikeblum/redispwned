package config

import (
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
)

const (
	envRedisAddr     = "REDIS_ADDR"
	defaultRedisAddr = "127.0.0.1:6379"
	envRedisPassword = "REDIS_PASSWORD"
	envRedisDatabase = "REDIS_DB"
	DefaultRedisDB   = 0
	TestRedisDB      = 1
)

func NewDefaultRedisClient() *redis.Client {
	var redisAddr string
	var redisPassword string
	var redisDB int
	var err error

	if redisDB, err = strconv.Atoi(GetEnv(envRedisDatabase, strconv.Itoa(DefaultRedisDB))); err != nil {
		redisDB = DefaultRedisDB
	}
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
}

func NewRedisClientTest() *redis.Client {
	var redisAddr string
	var redisDB int = TestRedisDB
	redisPassword := GetEnv(envRedisPassword, "")
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
}

func newRedisClient(redisDB int) *redis.Client {
	var redisAddr string
	redisPassword := GetEnv(envRedisPassword, "")
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
}

func NewRediSearchClient(indexName string) *redisearch.Client {
	redisAddr := GetEnv(envRedisAddr, defaultRedisAddr)
	return redisearch.NewClient(redisAddr, indexName)
}
