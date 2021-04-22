package geoip

import (
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
)

type Client struct {
	log *logrus.Entry
}

func NewClient() *Client {
	return &Client{
		log: config.NewLog(),
	}
}
