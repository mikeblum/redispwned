package index

import (
	"fmt"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/sirupsen/logrus"
)

const idxRedisVersionByCityCountryGeo = "idx:redis-version-by-country-city-geo"
const dropDocuments = false

type Manager struct {
	*redisearch.Client
	log *logrus.Entry
}

func NewManager(cfg *config.AppConfig) *Manager {
	rediSearch := config.NewRediSearchClientFromConfig(cfg, idxRedisVersionByCityCountryGeo)
	manager := &Manager{
		rediSearch,
		config.NewLog(),
	}
	return manager
}

func (idx *Manager) BuildIndex() error {
	// RediSearch will not index hashes whose fields do not match an existing index schema.
	// You can see the number of hashes not indexed using FT.INFO - hash_indexing_failures
	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("version", redisearch.TextFieldOptions{Sortable: true})).
		AddField(redisearch.NewTextField("ip")).
		AddField(redisearch.NewTextFieldOptions("city", redisearch.TextFieldOptions{Sortable: true})).
		AddField(redisearch.NewTextFieldOptions("country", redisearch.TextFieldOptions{Sortable: true}))
	_, err := idx.Info()
	if err == nil {
		err := idx.DropIndex(dropDocuments)
		if err != nil {
			return err
		}
	}

	// Create a index definition for automatic indexing on Hash updates.
	indexDefinition := redisearch.NewIndexDefinition().SetAsync(false).AddPrefix("censys:").AddPrefix("shodan:")

	// Add the Index Definition
	return idx.CreateIndexWithIndexDefinition(schema, indexDefinition)
}

func (idx *Manager) AwaitIndex() error {
	// Wait for all documents to be indexed
	var info *redisearch.IndexInfo
	var err error
	info, err = idx.Info()
	for info.IsIndexing {
		time.Sleep(time.Second)
		info, err = idx.Info()
	}

	if info.HashIndexingFailures > 0 {
		err = fmt.Errorf("[%s] failed to index %d documents - check for missing attributes", info.Name, info.HashIndexingFailures)
	}
	return err
}
