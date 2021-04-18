package config

import (
	"net/http"
	"strconv"
	"time"
)

const (
	envMaxConn            = "MAX_CONNECTIONS"
	defaultMaxConn        = 100
	envTimeoutSeconds     = "TIMEOUT_SECONDS"
	defaultTimeoutSeconds = 10
)

type ShodanClient struct {
	client *http.Client
}

func NewShodanClient() *ShodanClient {
	t := http.DefaultTransport.(*http.Transport).Clone()
	var err error
	if maxConn, err := strconv.Atoi(GetEnv(envMaxConn, strconv.Itoa(defaultMaxConn))); err == nil {
		t.MaxIdleConns = maxConn
		t.MaxConnsPerHost = maxConn
		t.MaxIdleConnsPerHost = maxConn
	}
	var timeoutSeconds int
	if timeoutSeconds, err = strconv.Atoi(GetEnv(envTimeoutSeconds, strconv.Itoa(defaultTimeoutSeconds))); err != nil {
		timeoutSeconds = defaultTimeoutSeconds
	}
	return &ShodanClient{
		client: &http.Client{
			Timeout:   time.Duration(timeoutSeconds) * time.Second,
			Transport: t,
		},
	}
}
