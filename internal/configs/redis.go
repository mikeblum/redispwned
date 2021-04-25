package config

import (
	"fmt"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
)

const (
	envRedisAddr     = "REDIS_ADDR"
	defaultRedisAddr = "127.0.0.1:6379"
	envRedisPassword = "REDIS_PASSWORD"
	envRedisDatabase = "REDIS_DB"
	envRedisPort     = "REDIS_PORT"
	DefaultRedisDB   = 0
	TestRedisDB      = 1
)

func NewRedisClientFromConfig(cfg *AppConfig) *redis.Client {
	redisAddr := cfg.GetString(envRedisAddr)
	redisPort := cfg.GetInt(envRedisPort)
	redisPassword := cfg.GetString(envRedisPassword)
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisAddr, redisPort),
		Password: redisPassword,
		DB:       0,
	})
}

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

func NewRediSearchClientFromConfig(cfg *AppConfig, idxName string) *redisearch.Client {
	pool := &redigo.Pool{Dial: func() (redigo.Conn, error) {
		return redigo.Dial("tcp", redisAddr(cfg), redigo.DialPassword(cfg.GetString(envRedisPassword)))
	}}
	client := redisearch.NewClientFromPool(pool, idxName)
	return client
}

func NewRediSearchClient(indexName string) *redisearch.Client {
	redisAddr := GetEnv(envRedisAddr, defaultRedisAddr)
	return redisearch.NewClient(redisAddr, indexName)
}

func redisAddr(cfg *AppConfig) string {
	return fmt.Sprintf("%s:%d", cfg.GetString(envRedisAddr), cfg.GetInt(envRedisPort))
}
