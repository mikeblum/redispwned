package geoip

import (
	"context"
	"net"
	"testing"

	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/redispwned/internal/configs"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suites.Suite
	client *Client
	redis  *redis.Client
	err    error
}

func (suite *ClientTestSuite) SetupSuite() {
	// use 1 in tests, 0 in production
	cfg := config.NewConfig()
	suite.redis = config.NewRedisClientTest(cfg)
	suite.client = NewClient(suite.redis)
	// !!DANGER!! flush db
	suite.redis.FlushDB(context.TODO())
	suite.err = suite.client.ImportGeoIPData()
}

func (suite *ClientTestSuite) TearDownAllSuite() {
	suite.redis.FlushDB(context.TODO())
}

func TestClientTestSuite(t *testing.T) {
	suites.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) TestLookupIP() {
	assert := asserts.New(suite.T())
	results, err := suite.client.LookupIP(net.ParseIP("8.8.8.8"))
	assert.Nil(err)
	assert.True(len(results) > 0)
}
