package account

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"

	censys "github.com/mikeblum/haveibeenredised/internal/censys"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	"github.com/sirupsen/logrus"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

const accountJSON = "test-data/account.json"

type AccountTestSuite struct {
	suites.Suite
	client *censys.Client
	cfg    *config.AppConfig
	log    *logrus.Entry
}

func (suite *AccountTestSuite) SetupTest() {
	suite.cfg = config.NewConfig()
	suite.client = censys.NewClient()
	suite.log = config.NewLog()
}

func TestAccountTestSuite(t *testing.T) {
	suites.Run(t, new(AccountTestSuite))
}

func (suite *AccountTestSuite) TestAccountConfig() {
	assert := asserts.New(suite.T())
	assert.NotNil(suite.cfg)
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
	account, err := GetAccount(suite.client)
	assert.Nil(err)
	assert.NotNil(account)

	data, err := json.Marshal(account)
	assert.Nil(err)
	suite.log.Info(string(data))
	// check quota and warn if depleted
	assert.True(account.Quota.Used < account.Quota.Allowance)
}
