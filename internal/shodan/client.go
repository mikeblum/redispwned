package shodan

import (
	"net/http"

	"github.com/RediSearch/redisearch-go/redisearch"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
)

type Client struct {
	client *http.Client
	idx    *redisearch.Client
	log    *logrus.Entry
}

func NewClient() *Client {
	return &Client{
		client: config.NewHTTPClient(),
		log:    config.NewLog(),
	}
}
