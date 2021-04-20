package geoip

import (
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
)

type GeoIPClient struct {
	log *logrus.Entry
}

func NewGeoIPClient() *GeoIPClient {
	return &GeoIPClient{
		log: config.NewLog(),
	}
}
