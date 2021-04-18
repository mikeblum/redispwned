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
	defaultRedisDB   = 0
)

func NewRedisClient() *redis.Client {
	var redisAddr string
	var redisPassword string
	var redisDB int
	var err error

	redisPassword = GetEnv(envRedisPassword, "")
	if redisDB, err = strconv.Atoi(GetEnv(envRedisDatabase, strconv.Itoa(defaultRedisDB))); err != nil {
		redisDB = defaultRedisDB
	}
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
