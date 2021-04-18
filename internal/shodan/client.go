package shodan

import (
	"net/http"

	"github.com/RediSearch/redisearch-go/redisearch"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
)

type ShodanClient struct {
	client *http.Client
	idx    *redisearch.Client
	log    *logrus.Entry
}

func NewShodanClient() *ShodanClient {
	return &ShodanClient{
		client: config.NewHTTPClient(),
		log:    config.NewLog(),
	}
}
