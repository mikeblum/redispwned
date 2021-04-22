package crawler

import (
	"encoding/json"
	"net"
	"testing"

	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
	"github.com/zmap/zgrab2"
)

type ScanTestSuite struct {
	suites.Suite
	scanner *RedisScanner
}

func (suite *ScanTestSuite) SetupTest() {
	var err error
	suite.scanner, err = NewScanner()
	assert := asserts.New(suite.T())
	assert.Nil(err)
}

func (suite *ScanTestSuite) TestNewScanner() {
	assert := asserts.New(suite.T())
	assert.NotNil(suite.scanner)
}

func (suite *ScanTestSuite) TestScanLocalhost() {
	assert := asserts.New(suite.T())
	localhost := net.ParseIP("127.0.0.1")
	assert.NotNil(localhost)
	target := zgrab2.ScanTarget{
		IP:   localhost,
		Port: suite.scanner.RedisPort(),
	}
	status, response, err := suite.scanner.Scan(target)
	assert.Nil(err)
	assert.Equal(zgrab2.SCAN_SUCCESS, status)
	data, err := json.Marshal(response)
	assert.Nil(err)
	assert.True(len(data) > 0)
}

func TestScanTestSuite(t *testing.T) {
	suites.Run(t, new(ScanTestSuite))
}
