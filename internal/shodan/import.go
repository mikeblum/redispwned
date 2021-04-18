package shodan

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
)

const DataJSONPath = "data/shodan-export.json"

const attrHashIndexingFailures = "hash_indexing_failures"

const idxShodanRedisVersionByCityCountryGeo = "idx:shodan-redis-version-by-country-city-geo"

func (s *ShodanClient) ImportShodanData(path string, redisClient *redis.Client) error {
	s.buildIndex()
	ctx := context.TODO()
	var err error
	dump, err := s.LoadFile(path)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(dump)
	decoder := json.NewDecoder(reader)
	numRecords := 0
	pipe := redisClient.TxPipeline()
	for {
		var meta RedisMeta
		if err = decoder.Decode(&meta); err == io.EOF {
			break
		} else if err != nil {
			s.log.Error("Error reading Shodan data: ", err)
			break
		}
		meta.ToHSet(ctx, pipe)
		numRecords++
	}
	results, _ := pipe.Exec(ctx)
	s.log.Infof("Loaded %d / %d Shodan records", numRecords, len(results))
	dump.Close()
	return s.awaitIndex()
}

func (s *ShodanClient) LoadFile(path string) (*os.File, error) {
	dump, err := os.Open(path)
	if err != nil {
		s.log.Error("Failed to load Shodan export data")
	}
	return dump, err
}

func (s *ShodanClient) buildIndex() error {
	s.idx = config.NewRediSearchClient(idxShodanRedisVersionByCityCountryGeo)

	// RediSearch will not index hashes whose fields do not match an existing index schema.
	// You can see the number of hashes not indexed using FT.INFO - hash_indexing_failures
	// shodan data is inconsistent especially with China data
	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("redis_version", redisearch.TextFieldOptions{Sortable: true})).
		AddField(redisearch.NewTextField("ip_addr")).
		// AddField(redisearch.NewTextField("isp")). missing for China data
		AddField(redisearch.NewTextFieldOptions("city", redisearch.TextFieldOptions{Sortable: true})).
		AddField(redisearch.NewTextFieldOptions("country", redisearch.TextFieldOptions{Sortable: true}))
	// AddField(redisearch.NewGeoField("geo")) missing for some entries
	s.idx.DropIndex(true)

	// IndexDefinition is available for RediSearch 2.0+
	// Create a index definition for automatic indexing on Hash updates.
	indexDefinition := redisearch.NewIndexDefinition().SetAsync(false).AddPrefix("shodan:")

	// Add the Index Definition
	return s.idx.CreateIndexWithIndexDefinition(schema, indexDefinition)

}

func (s *ShodanClient) awaitIndex() error {
	// Wait for all documents to be indexed
	var info *redisearch.IndexInfo
	var err error
	var hashIndexingFailures int
	info, err = s.idx.Info()
	for info.IsIndexing {
		time.Sleep(time.Second)
		info, err = s.idx.Info()
	}
	// hash_indexing_failures missing from *redisearch.IndexInfo
	rds := config.NewRedisClient()
	res, _ := rds.Do(context.TODO(), "FT.INFO", idxShodanRedisVersionByCityCountryGeo).Result()
	props := res.([]interface{})
	for i, itr := range props {
		prop, ok := itr.(string)
		if ok && strings.EqualFold(prop, attrHashIndexingFailures) {
			hashIndexingFailures, err = strconv.Atoi(props[i+1].(string))
			break
		}
	}
	if hashIndexingFailures > 0 {
		err = fmt.Errorf("failed to index %d documents - check for mssing indexed attributes", hashIndexingFailures)
	}
	return err
}
