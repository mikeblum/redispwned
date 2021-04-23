package censys

import (
	"testing"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

type IndexTestSuite struct {
	suites.Suite
	client *Client
	cfg    *config.AppConfig
	log    *logrus.Entry
}

func (suite *IndexTestSuite) SetupTest() {
	suite.cfg = config.NewConfig()
	suite.client = NewClient()
	suite.log = config.NewLog()
}

func TestIndexTestSuite(t *testing.T) {
	suites.Run(t, new(IndexTestSuite))
}

func (suite *IndexTestSuite) TestBuildIndex() {
	assert := asserts.New(suite.T())
	err := suite.client.buildIndex()
	assert.Nil(err)
}

func (suite *IndexTestSuite) TestAwaitIndex() {
	assert := asserts.New(suite.T())
	err := suite.client.awaitIndex()
	assert.Nil(err)
}
