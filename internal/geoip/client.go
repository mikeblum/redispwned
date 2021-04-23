package geoip

import (
	config "github.com/mikeblum/redispwned/internal/configs"
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
