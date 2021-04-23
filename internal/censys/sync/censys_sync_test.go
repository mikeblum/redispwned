// +build sync

package sync

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/mikeblum/redispwned/internal/censys"
	"github.com/mikeblum/redispwned/internal/censys/search"
	asserts "github.com/stretchr/testify/assert"
)

const censysDataJSONPath = "/tmp/censys-redis.json"

// Rate Limits
// web: 0.2 actions/second (60.0 per 5 minute interval)
// api: 0.4 actions/second (120.0 per 5 minute interval)
const rateLimitSeconds = 2

// !! WARNING - Censys API Quotas Allow For Only 2X / Month Syncs !!
// Allowed Queries: 250 x Results per search query Up to 1,000
// ~90k results --> 90 queries
func TestCensysRedisDataExport(t *testing.T) {
	assert := asserts.New(t)
	data, err := os.Open(censysDataJSONPath)
	assert.Nil(err)
	encoder := json.NewEncoder(data)
	client := censys.NewClient()
	page := 1 // Censys API is one-indexed
	for {
		response, err := search.RedisQuery(client, page)
		assert.Nil(err)
		assert.NotNil(response)
		client.Log.Infof("[ Page # %d / %d ] %d / %d results", response.Meta.Page, response.Meta.NumPages, len(response.Results), response.Meta.Count)
		for _, result := range response.Results {
			encoder.Encode(result)
		}
		page++
		if page > response.Meta.NumPages {
			break
		}
		time.Sleep(rateLimitSeconds * time.Second)
		break
	}
	data.Close()
}
