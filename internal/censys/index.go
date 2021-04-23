package censys

import (
	"fmt"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
)

const idxName = "censys:"
const idxRedisVersionByCityCountryGeo = "idx:redis-version-by-country-city-geo"

func (c *Client) buildIndex() error {
	c.Idx = config.NewRediSearchClient(idxRedisVersionByCityCountryGeo)

	// RediSearch will not index hashes whose fields do not match an existing index schema.
	// You can see the number of hashes not indexed using FT.INFO - hash_indexing_failures
	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("redis_version", redisearch.TextFieldOptions{Sortable: true})).
		AddField(redisearch.NewTextField("ip")).
		// AddField(redisearch.NewTextField("isp")). missing for China data
		AddField(redisearch.NewTextFieldOptions("city", redisearch.TextFieldOptions{Sortable: true})).
		AddField(redisearch.NewTextFieldOptions("country", redisearch.TextFieldOptions{Sortable: true}))
	// AddField(redisearch.NewGeoField("geo")) missing for some entries
	_, err := c.Idx.Info()
	if err == nil {
		err := c.Idx.DropIndex(true)
		if err != nil {
			return nil
		}
	}

	// IndexDefinition is available for RediSearch 2.0+
	// Create a index definition for automatic indexing on Hash updates.
	indexDefinition := redisearch.NewIndexDefinition().SetAsync(false).AddPrefix(idxName)

	// Add the Index Definition
	return c.Idx.CreateIndexWithIndexDefinition(schema, indexDefinition)
}

func (c *Client) awaitIndex() error {
	// Wait for all documents to be indexed
	var info *redisearch.IndexInfo
	var err error
	info, err = c.Idx.Info()
	for info.IsIndexing {
		time.Sleep(time.Second)
		info, err = c.Idx.Info()
	}

	if info.HashIndexingFailures > 0 {
		err = fmt.Errorf("[%s] failed to index %d documents - check for missing attributes", info.Name, info.HashIndexingFailures)
	}
	return err
}
