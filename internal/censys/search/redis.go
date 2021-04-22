package search

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mikeblum/haveibeenredised/internal/censys"
)

const endpointSearchIPV4 = "search/ipv4"

const queryRedis = `protocols: "6379/redis" OR tags.raw: "redis"`

func Redis(c *censys.Client) (*Response, error) {
	var response Response
	var err error
	var req *http.Request
	var res *http.Response
	ctx := context.Background()
	search := Request{
		Query:   queryRedis,
		Page:    1,
		Fields:  make([]string, 0),
		Flatten: false,
	}
	payload, err := json.Marshal(search)
	c.Log.Info(string(payload))
	if err != nil {
		return &response, err
	}
	req, err = c.NewAPIRequest(ctx, http.MethodPost, endpointSearchIPV4, bytes.NewBuffer(payload))
	if err != nil {
		return &response, err
	}
	res, err = c.API.Do(req)
	apiErr := c.Err(res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	decoder := json.NewDecoder(res.Body)
	if decodeErr := decoder.Decode(&response); decodeErr != nil {
		return &response, decodeErr
	}
	data, err := ioutil.ReadAll(res.Body)
	c.Log.Info(string(data))
	res.Body.Close()
	return &response, err
}
