package api

import (
	"context"

	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/mikeblum/redispwned/internal/index"
)

func NewRedisConn() (*redis.Client, context.Context) {
	ctx := context.TODO()
	cfg := config.NewConfig()
	return config.NewRedisClientFromConfig(cfg), ctx
}

func NewSearchEngine() *index.Manager {
	cfg := config.NewConfig()
	return index.NewManager(cfg)
}
