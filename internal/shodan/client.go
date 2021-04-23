package shodan

import (
	"net/http"

	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/sirupsen/logrus"
)

type Client struct {
	client *http.Client
	log    *logrus.Entry
}

func NewClient() *Client {
	return &Client{
		client: config.NewHTTPClient(),
		log:    config.NewLog(),
	}
}
