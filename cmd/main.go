package main

import (
	"context"

	"github.com/go-redis/redis/v8"

	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/mikeblum/redispwned/internal/geoip"
	"github.com/mikeblum/redispwned/internal/index"
	"github.com/mikeblum/redispwned/internal/shodan"
)

var ctx = context.Background()

func main() {
	var err error
	log := config.NewLog()
	cfg := config.NewConfig()
	redisClient := config.NewRedisClientFromConfig(cfg)
	log.Info(redisClient.Ping(ctx))
	err = loadGeoIPData(redisClient)
	if err != nil {
		log.Fatal("Failed to load GeoIP data: ", err)
	}
	err = loadShodanData(redisClient)
	if err != nil {
		log.Fatal("Failed to load Shodan data: ", err)
	}
	idx := index.NewManager(cfg)
	err = buildIndexes(idx)
	if err != nil {
		log.Fatal("Failed to build search indexes: ", err)
	}
	err = idx.ServersByCountry()
	if err != nil {
		log.Warn("servers by country query failed: ", err)
	}
	err = idx.ServersByVersion()
	if err != nil {
		log.Warn("servers by version query failed: ", err)
	}
}

func loadGeoIPData(redisClient *redis.Client) error {
	geoIPClient := geoip.NewClient()
	return geoIPClient.ImportGeoIPData(redisClient)
}

func loadShodanData(redisClient *redis.Client) error {
	shodanClient := shodan.NewClient()
	return shodanClient.ImportShodanData(redisClient)
}

func buildIndexes(idx *index.Manager) error {
	err := idx.BuildIndex()
	if err != nil {
		return err
	}
	return idx.AwaitIndex()
}
