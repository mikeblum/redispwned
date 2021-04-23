package account

import (
	"context"
	"encoding/json"
	"net/http"

	censys "github.com/mikeblum/redispwned/internal/censys"
)

const endpointAccount = "account"

func GetAccount(c *censys.Client) (*Account, error) {
	var account Account
	var err error
	var req *http.Request
	var res *http.Response
	ctx := context.Background()
	req, err = c.NewAPIRequest(ctx, http.MethodGet, endpointAccount, nil)
	if err != nil {
		return nil, err
	}
	res, err = c.API.Do(req)
	apiErr := c.Err(res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}
	decoder := json.NewDecoder(res.Body)
	if decodeErr := decoder.Decode(&account); decodeErr != nil {
		return nil, decodeErr
	}
	res.Body.Close()
	return &account, err
}
