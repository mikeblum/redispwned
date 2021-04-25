package index

import (
	"context"
	"testing"

	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/sirupsen/logrus"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

type IndexTestSuite struct {
	suites.Suite
	idx *Manager
	cfg *config.AppConfig
	log *logrus.Entry
}

func (suite *IndexTestSuite) SetupTest() {
	suite.cfg = config.NewConfig()
	suite.idx = NewManager(suite.cfg)
	suite.log = config.NewLog()
}

func (suite *IndexTestSuite) SetupSuite() {
	// use 1 in tests, 0 in production
	client := config.NewRedisClientTest()
	// !!DANGER!! flush db before each test
	client.FlushDB(context.TODO())
}

func TestIndexTestSuite(t *testing.T) {
	suites.Run(t, new(IndexTestSuite))
}

func (suite *IndexTestSuite) TestBuildIndex() {
	assert := asserts.New(suite.T())
	var err error
	err = suite.idx.BuildIndex()
	assert.Nil(err)
	err = suite.idx.AwaitIndex()
	assert.Nil(err)
}
