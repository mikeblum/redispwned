package main

import (
	"context"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/mikeblum/haveibeenredised/internal/shodan"
)

var ctx = context.Background()

func main() {
	log := config.NewLog()
	redisClient := config.NewRedisClient()
	log.Info(redisClient.Ping(ctx))
	shodanClient := shodan.NewShodanClient()
	shodanClient.ImportShodanData(shodan.DataJSONPath, redisClient)
	shodanClient.ServersByCountry()
	shodanClient.ServersByVersion()
}
