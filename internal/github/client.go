package github

import (
	"net/http"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
)

type GithubClient struct {
	client *http.Client
	log    *logrus.Entry
}

func NewGithubClient() *GithubClient {
	return &GithubClient{
		client: config.NewHTTPClient(),
		log:    config.NewLog(),
	}
}
