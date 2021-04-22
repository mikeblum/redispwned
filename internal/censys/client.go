package censys

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
)

const envCensysID = "CENSYS_ID"
const envCensysKey = "CENSYS_SECRET"
const apiURL = "https://censys.io/api/v1"

const applicationJSON = "application/json"
const headerContentType = "Content-Type"
const headerAccept = "Accept"

type Client struct {
	API *http.Client
	Cfg *config.AppConfig
	Log *logrus.Entry
}

func NewClient() *Client {
	return &Client{
		API: config.NewHTTPClient(),
		Cfg: config.NewConfig(),
		Log: config.NewLog(),
	}
}

func (c *Client) NewAPIRequest(ctx context.Context, method, path string, payload io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, strings.Join([]string{apiURL, path}, "/"), payload)
	c.Log.Debug(req.URL.String())
	if err != nil {
		return nil, err
	}
	// headers
	req.Header.Add(headerContentType, applicationJSON)
	req.Header.Add(headerAccept, applicationJSON)
	req = c.Auth(req)
	return req, err
}

func (c *Client) Auth(req *http.Request) *http.Request {
	ID := c.Cfg.GetString(envCensysID)
	key := c.Cfg.GetString(envCensysKey)
	req.SetBasicAuth(ID, key)
	return req
}

func (c *Client) Err(res *http.Response) error {
	// err > 400 Bad Request
	if res.StatusCode > http.StatusBadRequest {
		return fmt.Errorf(res.Status)
	}
	return nil
}
