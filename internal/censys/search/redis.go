package search

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/mikeblum/redispwned/internal/censys"
)

const endpointSearchIPV4 = "search/ipv4"

const queryRedis = `location.country_code: US and protocols: "6379/redis"`

// RedisQuery - page: API will return the first page of results. One indexed.
func RedisQuery(c *censys.Client, page int) (*Response, error) {
	var response Response
	var err error
	var req *http.Request
	var res *http.Response
	ctx := context.Background()
	fields := make([]string, 0)
	// concat censys fields
	fields = append(fields, redisFields()...)
	fields = append(fields, asnFields()...)
	fields = append(fields, ipFields()...)
	fields = append(fields, locationFields()...)
	fields = append(fields, metaFields()...)
	search := Request{
		Query:   queryRedis,
		Page:    page,
		Fields:  fields,
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
		return &response, err
	} else if apiErr != nil {
		return &response, apiErr
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if decodeErr := decoder.Decode(&response); decodeErr != nil {
		return &response, decodeErr
	}
	res.Body.Close()
	return &response, err
}

// START HERE: https://censys.io/ipv4/help/definitions

func redisFields() []string {
	return []string{
		// this exposes too much info
		// "6379.redis.banner.info_response",
		"6379.redis.banner.ping_response",
		"6379.redis.banner.version",
		"6379.redis.banner.mode",
		"6379.redis.banner.uptime_in_seconds",
	}
}

func asnFields() []string {
	return []string{
		"autonomous_system.asn",
		"autonomous_system.country_code", // ASN two-letter ISO 3166-1 (US, CN, GB, RU, ...).
		"autonomous_system.description",
		"autonomous_system.name",
		"autonomous_system.organization",
		"autonomous_system.routed_prefix", // The autonomous system's CIDR.
	}
}

func ipFields() []string {
	return []string{
		"ip",
	}
}

func locationFields() []string {
	return []string{
		"location.city",
		"location.province",
		"location.country",
		"location.country_code",
		"location.latitude",
		"location.longitude",
	}
}

func metaFields() []string {
	return []string{
		"updated_at",
	}
}
