package censys

import (
	"encoding/json"
	"net/http"
)

const endpointAccount = "/account"

func (c *Client) GetAccount() (*Account, error) {
	var account Account
	var err error
	var req *http.Request
	var res *http.Response
	req, err = c.NewRequest(endpointAccount)
	if err != nil {
		return nil, err
	}
	res, err = c.api.Do(req)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	if decodeErr := decoder.Decode(&account); decodeErr != nil {
		return nil, decodeErr
	}
	res.Body.Close()
	return &account, err
}
