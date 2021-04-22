package main

import (
	"context"

	"github.com/go-redis/redis/v8"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/mikeblum/haveibeenredised/internal/geoip"
)

var ctx = context.Background()

func main() {
	log := config.NewLog()
	redisClient := config.NewDefaultRedisClient()
	log.Info(redisClient.Ping(ctx))
	err := loadGeoIPData(redisClient)
	if err != nil {
		log.Fatal("Failed to load GeoIP data: ", err)
	}
}

func loadGeoIPData(redisClient *redis.Client) error {
	geoIPClient := geoip.NewClient()
	return geoIPClient.ImportGeoIPData(redisClient)
}
