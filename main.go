package main

import (
	"context"

	config "github.com/mikeblum/haveibeenredised/pkg"
)

var ctx = context.Background()

func main() {
	log := config.NewLog()
	_ = config.NewConfig()
	redisClient := config.NewRedisClient()
	log.Info(redisClient.Ping(ctx))
}
