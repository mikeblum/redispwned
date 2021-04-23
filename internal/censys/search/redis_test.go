package search

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"

	censys "github.com/mikeblum/redispwned/internal/censys"
	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/sirupsen/logrus"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

const searchResultsJSON = "test-data/results.json"

type SearchTestSuite struct {
	suites.Suite
	client *censys.Client
	cfg    *config.AppConfig
	log    *logrus.Entry
}

func (suite *SearchTestSuite) SetupTest() {
	suite.cfg = config.NewConfig()
	suite.client = censys.NewClient()
	suite.log = config.NewLog()
}

func TestSearchTestSuite(t *testing.T) {
	suites.Run(t, new(SearchTestSuite))
}

func (suite *SearchTestSuite) TestSearchConfig() {
	assert := asserts.New(suite.T())
	assert.NotNil(suite.cfg)
}

func (suite *SearchTestSuite) TestSearchResultSerialization() error {
	assert := asserts.New(suite.T())
	var response Response
	dump, err := os.Open(searchResultsJSON)
	if err != nil {
		suite.log.Error("Failed to load search results")
		return err
	}
	reader := bufio.NewReader(dump)
	decoder := json.NewDecoder(reader)
	for {
		if err = decoder.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			suite.log.Error("Error reading search results: ", err)
			assert.Nil(err)
			break
		}
	}
	assert.NotNil(response)
	expectedResults := 3
	assert.Equal(expectedResults, len(response.Results))
	expectedTotalPages := 910
	assert.Equal(expectedTotalPages, response.Meta.NumPages)
	expectedTotalResults := 90905
	assert.Equal(expectedTotalResults, response.Meta.Count)
	return nil
}

func mockResponse() Response {
	return Response{
		Results: []Result{
			{
				Redis: Redis{
					Service: Service{
						Banner: Banner{
							PingResponse: "PONG",
						},
					},
				},
			},
		},
	}
}

func (suite *SearchTestSuite) TestPingResponsePONG() {
	assert := asserts.New(suite.T())
	response := mockResponse()
	assert.NotNil(response)
}

func (suite *SearchTestSuite) TestPingResponseNOAUTH() {
	assert := asserts.New(suite.T())
	response := mockResponse()
	response.Results[0].Redis.Service.Banner.PingResponse = "(Error: NOAUTH Authentication required.)"
	assert.NotNil(response)
}
