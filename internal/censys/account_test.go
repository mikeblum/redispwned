package censys

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"

	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

const accountJSON = "test-data/account.json"

type AccountTestSuite struct {
	suites.Suite
	client *Client
	log    *logrus.Entry
}

func (suite *AccountTestSuite) SetupTest() {
	suite.client = NewClient()
	suite.log = config.NewLog()
}

func TestAccountTestSuite(t *testing.T) {
	suites.Run(t, new(AccountTestSuite))
}

func (suite *AccountTestSuite) TestAccountSerialization() error {
	assert := asserts.New(suite.T())
	var account Account
	dump, err := os.Open(accountJSON)
	if err != nil {
		suite.log.Error("Failed to load account data")
		return err
	}
	reader := bufio.NewReader(dump)
	decoder := json.NewDecoder(reader)
	for {
		if err = decoder.Decode(&account); err == io.EOF {
			break
		} else if err != nil {
			suite.log.Error("Error reading account data: ", err)
			break
		}
	}
	assert.NotNil(account)
	return nil
}

func (suite *AccountTestSuite) TestAccountRequest() {
	assert := asserts.New(suite.T())
	account, err := suite.client.GetAccount()
	assert.Nil(err)
	assert.NotNil(account)
}
