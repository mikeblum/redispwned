package shodan

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

const shodanDataJSONTest = "test-data/shodan.json"

type ImportTestSuite struct {
	suites.Suite
	client *redis.Client
	err    error
}

func (suite *ImportTestSuite) SetupTest() {
	// use 1 in tests, 0 in production
	suite.client = config.NewRedisClientTest()
	// !!DANGER!! flush db before each test
	response := suite.client.FlushDB(context.TODO())
	suite.err = response.Err()
}

func (suite *ImportTestSuite) TearDownAllSuite() {
	suite.client.FlushDB(context.TODO())
}

func TestImportTestSuite(t *testing.T) {
	suites.Run(t, new(ImportTestSuite))
}

func (suite *ImportTestSuite) TestImportShodanDataSerialization() {
	assert := asserts.New(suite.T())
	s := NewClient()
	assert.NotNil(s)
	err := s.ImportShodanData(shodanDataJSONTest, suite.client)
	assert.Nil(err)
}

// https://docs.redislabs.com/latest/modules/redisearch/release-notes/redisearch-2.0-release-notes/
func (suite *ImportTestSuite) TestImportShodanDataHashIndexingFailures() {
	// RediSearch will not index hashes whose fields do not match an existing index schema.
	// You can see the number of hashes not indexed using FT.INFO - hash_indexing_failures.
	// The requirement for adding support for partially indexing and blocking is captured here: #1455
	assert := asserts.New(suite.T())
	s := NewClient()
	redisClient := config.NewRedisClientTest()
	assert.NotNil(s)
	err := s.ImportShodanData("../../"+DataJSONPath, redisClient)
	assert.Nil(err)
}
