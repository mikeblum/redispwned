package censys

import (
	"net/http"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
)

const envCensysID = "CENSYS_ID"
const envCensysKey = "CENSYS_KEY"
const apiURL = "https://censys.io/api/v1"

const applicationJSON = "application/json"
const headerContentType = "Content-Type"
const headerAccept = "Accept"

type Client struct {
	api *http.Client
	log *logrus.Entry
}

func NewClient() *Client {
	return &Client{
		api: config.NewHTTPClient(),
		log: config.NewLog(),
	}
}

func (c *Client) NewRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	// headers
	req.Header.Add(headerContentType, applicationJSON)
	req.Header.Add(headerAccept, applicationJSON)
	req = c.Auth(req)
	return req, err
}

func (c *Client) Auth(req *http.Request) *http.Request {
	ID := config.GetEnv(envCensysID, "")
	key := config.GetEnv(envCensysKey, "")
	req.SetBasicAuth(ID, key)
	return req
}
