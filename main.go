package main

import (
	"context"

	"github.com/go-redis/redis/v8"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/mikeblum/haveibeenredised/internal/geoip"
	"github.com/mikeblum/haveibeenredised/internal/shodan"
)

var ctx = context.Background()

func main() {
	log := config.NewLog()
	redisClient := config.NewDefaultRedisClient()
	log.Info(redisClient.Ping(ctx))
	loadGeoIPData(redisClient)
}

func loadGeoIPData(redisClient *redis.Client) {
	geoIPClient := geoip.NewGeoIPClient()
	geoIPClient.ImportGeoIPData(redisClient)
}

func loadShodanData(redisClient *redis.Client) {
	shodanClient := shodan.NewShodanClient()
	shodanClient.ImportShodanData(shodan.DataJSONPath, redisClient)
	shodanClient.ServersByCountry()
	shodanClient.ServersByVersion()
}
