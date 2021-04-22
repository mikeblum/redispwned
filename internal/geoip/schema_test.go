package geoip

import (
	"context"
	"net"
	"testing"

	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/haveibeenredised/internal/configs"
	asserts "github.com/stretchr/testify/assert"
	suites "github.com/stretchr/testify/suite"
)

type GeoIPTestSuite struct {
	suites.Suite
	client *redis.Client
	err    error
}

func (suite *GeoIPTestSuite) SetupTest() {
	// use 1 in tests, 0 in production
	suite.client = config.NewRedisClientTest()
	// !!DANGER!! flush db before each test
	response := suite.client.FlushDB(context.TODO())
	suite.err = response.Err()
}

func (suite *GeoIPTestSuite) TearDownAllSuite() {
	suite.client.FlushDB(context.TODO())
}

func TestGeoIPTestSuite(t *testing.T) {
	suites.Run(t, new(GeoIPTestSuite))
}

func (suite *GeoIPTestSuite) TestImportGeoCIDRToInt() {
	assert := asserts.New(suite.T())
	googleDNS := "8.8.8.8/18"
	broadcast, network, err := net.ParseCIDR(googleDNS)
	assert.Nil(err)
	assert.NotNil(network)
	// calculated using python netaddr
	// from netaddr import IPNetwork
	// >>> cidr = '8.8.0.0/18'
	// >>> network = int(IPNetwork(cidr).network)
	// >>> broadcast = int(IPNetwork(cidr).broadcast)
	const expectedBroadcast int64 = 134744072
	const expectedNetwork int64 = 134742016
	broadcastNum, err := IPV4ToInt(broadcast)
	assert.Nil(err)
	assert.Equal(expectedBroadcast, broadcastNum)
	networkNum, err := IPV4ToInt(network.IP)
	assert.Nil(err)
	assert.Equal(expectedNetwork, networkNum)
}

func (suite *GeoIPTestSuite) TestImportGeoIPV4ToInt() {
	assert := asserts.New(suite.T())
	ipv4 := "172.31.40.236"
	ipAddr := net.ParseIP(ipv4)
	// private AWS EC2 IP
	// >>> int(IPAddress('172.31.40.236'))
	var expectedAddrNum int64 = 2887723244
	addrNum, err := IPV4ToInt(ipAddr)
	assert.Nil(err)
	assert.Equal(expectedAddrNum, addrNum)
	// public AWS EC2 IP
	// >>> int(IPAddress('3.134.110.210'))
	ipv4 = "3.134.110.210"
	ipAddr = net.ParseIP(ipv4)
	expectedAddrNum = 59141842
	addrNum, err = IPV4ToInt(ipAddr)
	assert.Nil(err)
	assert.Equal(expectedAddrNum, addrNum)
}

func (suite *GeoIPTestSuite) TestImportGeoIpValidCoordinates() {
	assert := asserts.New(suite.T())
	var lon float64
	var lat float64
	var err error
	// we're not in Kansas anymore Toto
	lat, err = CoordinateToFloat("39.8333333")
	assert.Nil(err)
	assert.True(ValidLat(lat))
	lon, err = CoordinateToFloat("-98.585522")
	assert.Nil(err)
	assert.True(ValidLon(lon))
}

func (suite *GeoIPTestSuite) TestImportGeoIpInvalidCoordinates() {
	assert := asserts.New(suite.T())
	// Failed to import GeoIP data ERR invalid longitude,latitude pair 6252001.000000,1861060.000000
	var lon float64 = 6252001.000000
	var lat float64 = 1861060.000000
	assert.False(ValidLon(lon))
	assert.False(ValidLat(lat))
}
