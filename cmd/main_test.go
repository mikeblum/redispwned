package main

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/mikeblum/redispwned/internal/index"
	"github.com/sirupsen/logrus"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suites.Suite
	client *redis.Client
	cfg    *config.AppConfig
	log    *logrus.Entry
}

func (suite *MainTestSuite) SetupSuite() {
	// use 1 in tests, 0 in production
	cfg := config.NewConfig()
	suite.client = config.NewRedisClientTest(cfg)
	suite.cfg = config.NewConfig()
	suite.log = config.NewLog()
	// !!DANGER!! flush db before each test
	suite.client.FlushDB(context.TODO())
}

func TestMainTestSuite(t *testing.T) {
	suites.Run(t, new(MainTestSuite))
}

func (suite *MainTestSuite) TestLoadGeoIPData() {
	assert := asserts.New(suite.T())
	err := loadGeoIPData(suite.client)
	assert.Nil(err)
}

func (suite *MainTestSuite) TestLoadShodanData() {
	assert := asserts.New(suite.T())
	err := loadShodanData(suite.client)
	assert.Nil(err)
}

func (suite *MainTestSuite) TestBuildIndexes() {
	assert := asserts.New(suite.T())
	mgr := index.NewManager(suite.cfg)
	err := buildIndexes(mgr)
	assert.Nil(err)
}
