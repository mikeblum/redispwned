package shodan

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const shodanDataJSONTest = "test-data/shodan.json"

type ImportTestSuite struct {
	suite.Suite
	client *redis.Client
	err    error
}

func (suite *ImportTestSuite) SetupTest() {
	suite.client = config.NewRedisClient()
	// !!DANGER!! flush db before each test
	response := suite.client.FlushDB(context.TODO())
	suite.err = response.Err()
}

func TestImportTestSuite(t *testing.T) {
	suite.Run(t, new(ImportTestSuite))
}

func (suite *ImportTestSuite) TestImportShodanDataSerialization() {
	assert := assert.New(suite.T())
	s := NewShodanClient()
	redisClient := config.NewRedisClient()
	assert.NotNil(s)
	err := s.ImportShodanData(shodanDataJSONTest, redisClient)
	assert.Nil(err)
}

// https://docs.redislabs.com/latest/modules/redisearch/release-notes/redisearch-2.0-release-notes/
func (suite *ImportTestSuite) TestImportShodanDataHashIndexingFailures() {
	// RediSearch will not index hashes whose fields do not match an existing index schema.
	// You can see the number of hashes not indexed using FT.INFO - hash_indexing_failures.
	// The requirement for adding support for partially indexing and blocking is captured here: #1455
	assert := assert.New(suite.T())
	s := NewShodanClient()
	redisClient := config.NewRedisClient()
	assert.NotNil(s)
	err := s.ImportShodanData("../../"+DataJSONPath, redisClient)
	assert.Nil(err)
}
