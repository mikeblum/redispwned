package crawler

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zmap/zgrab2"
)

type ScanTestSuite struct {
	suite.Suite
	scanner *RedisScanner
}

func (suite *ScanTestSuite) SetupTest() {
	suite.scanner = NewScanner()
}

func (suite *ScanTestSuite) TestNewScanner() {
	assert := assert.New(suite.T())
	assert.NotNil(suite.scanner)
}

func (suite *ScanTestSuite) TestScanLocalhost() {
	// assert := assert.New(suite.T())
	target := &zgrab2.ScanTarget{
		IP:   net.ParseIP("127.0.0.1"),
		Port: suite.scanner.RedisPort(),
	}
	suite.scanner.Scan(*target)
}

func TestScanTestSuite(t *testing.T) {
	suite.Run(t, new(ScanTestSuite))
}
