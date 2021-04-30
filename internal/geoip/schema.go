package geoip

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"net"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

const asnRedisKey = "asn"
const cidrRedisKey = "cidr"
const geoRedisKey = "geo"

type CityLocation struct {
	GeonameID           int    `csv:"geoname_id"`
	LocaleCode          string `csv:"-"` // - means unused
	ContinentCode       string `csv:"-"`
	ContinentName       string `csv:"-"`
	CountryISOCode      string `csv:"country_iso_code"`
	CountryName         string `csv:"country_name"`
	Subdivision1ISOCode string `csv:"-"`
	Subdivision1Name    string `csv:"-"`
	Subdivision2ISOCode string `csv:"-"`
	Subdivision2Name    string `csv:"-"`
	CityName            string `csv:"city_name"`
	MetroCode           string `csv:"-"`
	TZ                  string `csv:"time_zone"`
	IsEU                bool   `csv:"is_in_european_union"`
}

func NewCityLocation(row []string) interface{} {
	geonameID, _ := strconv.Atoi(row[0])
	isEu, _ := strconv.Atoi(row[6])
	return &CityLocation{
		GeonameID:      geonameID,
		CountryISOCode: row[4],
		CountryName:    row[5],
		CityName:       row[10],
		TZ:             row[12],
		IsEU:           isEu != 0,
	}
}

func (cl *CityLocation) GetRedisKey() string {
	geoname := strconv.Itoa(cl.GeonameID)
	return strings.Join([]string{geoRedisKey, geoname}, ":")
}

func (cl *CityLocation) ToHSet(ctx context.Context, pipe redis.Pipeliner) {
	pipe.HSet(ctx, cl.GetRedisKey(),
		"country_code", cl.CountryISOCode,
		"city", cl.CityName,
		"tz", cl.TZ)
}

type CityBlock struct {
	CIDR      string
	Broadcast net.IP
	Network   *net.IPNet `csv:"network"`
	GeonameID int        `csv:"geoname_id"`
	Latitude  float64    `csv:"latitude"`
	Longitude float64    `csv:"longitude"`
}

func NewCityBlock(row []string) interface{} {
	geonameID, _ := strconv.Atoi(row[1])
	lat, _ := CoordinateToFloat(row[2])
	lon, _ := CoordinateToFloat(row[3])
	broadcast, network, err := net.ParseCIDR(row[0])
	if err != nil {
		fmt.Println("Failed to resolve CIDR: " + row[0])
	}
	return &CityBlock{
		CIDR:      row[0],
		Broadcast: broadcast,
		Network:   network,
		GeonameID: geonameID,
		Latitude:  lat,
		Longitude: lon,
	}
}

func (cb *CityBlock) GetRedisKey() string {
	geoname := strconv.Itoa(cb.GeonameID)
	return strings.Join([]string{geoRedisKey, geoname}, ":")
}

func (cb *CityBlock) ToHSet(ctx context.Context, pipe redis.Pipeliner) {
	if broadcast, err := IPV4ToInt(cb.Broadcast); err == nil {
		pipe.ZAdd(ctx, cidrRedisKey, &redis.Z{
			Score:  float64(broadcast),
			Member: cb.GeonameID,
		})
	}
	// The command takes arguments in the standard format x,y
	// so the longitude must be specified before the latitude
	// if ValidLon(cb.Longitude) && ValidLat(cb.Latitude) {
	// 	pipe.GeoAdd(ctx, cb.GetRedisKey(), &redis.GeoLocation{
	// 		Longitude: cb.Longitude,
	// 		Latitude:  cb.Latitude,
	// 	})
	// }
}

type CountryLocation struct {
	GeonameID      int    `csv:"geoname_id"`
	LocaleCode     string `csv:"-"`
	ContinentCode  string `csv:"-"`
	ContinentName  string `csv:"-"`
	CountryISOCode string `csv:"country_iso_code"`
	CountryName    string `csv:"country_name"`
	IsEU           bool   `csv:"is_in_european_union"`
}

func NewCountryLocation(row []string) interface{} {
	geonameID, _ := strconv.Atoi(row[0])
	isEu, _ := strconv.Atoi(row[6])
	return &CountryLocation{
		GeonameID:      geonameID,
		CountryISOCode: row[4],
		CountryName:    row[5],
		IsEU:           isEu != 0,
	}
}

func (cl *CountryLocation) GetRedisKey() string {
	geoname := strconv.Itoa(cl.GeonameID)
	return strings.Join([]string{geoRedisKey, geoname}, ":")
}

func (cl *CountryLocation) ToHSet(ctx context.Context, pipe redis.Pipeliner) {
	pipe.HSet(ctx, cl.GetRedisKey(),
		"country_code", cl.CountryISOCode,
		"country", cl.CountryName)
}

type CountryBlock struct {
	CIDR      string
	Broadcast net.IP
	Network   *net.IPNet `csv:"network"`
	GeonameID int        `csv:"geoname_id"`
}

func NewCountryBlock(row []string) interface{} {
	broadcast, network, _ := net.ParseCIDR(row[0])
	geonameID, _ := strconv.Atoi(row[1])
	return &CountryBlock{
		CIDR:      row[0],
		Broadcast: broadcast,
		Network:   network,
		GeonameID: geonameID,
	}
}

func (cb *CountryBlock) ToHSet(ctx context.Context, pipe redis.Pipeliner) {
	if broadcast, err := IPV4ToInt(cb.Broadcast); err == nil {
		pipe.ZAdd(ctx, cidrRedisKey, &redis.Z{
			Score:  float64(broadcast),
			Member: cb.GeonameID,
		})
	}
}

type ASNBlock struct {
	CIDR      string
	Broadcast net.IP
	Network   *net.IPNet `csv:"network"`
	ASN       int        `csv:"autonomous_system_number"`
	Org       string     `csv:"autonomous_system_organization"`
}

func NewASNBlock(row []string) interface{} {
	broadcast, network, _ := net.ParseCIDR(row[0])
	asn, _ := strconv.Atoi(row[1])
	return &ASNBlock{
		CIDR:      row[0],
		Broadcast: broadcast,
		Network:   network,
		ASN:       asn,
		Org:       row[2],
	}
}

// find matching ASN for a CIDRspace if any
func (asn *ASNBlock) GetRedisKey() string {
	return strings.Join([]string{asnRedisKey, asn.CIDR}, ":")
}

func (asn *ASNBlock) ToHSet(ctx context.Context, pipe redis.Pipeliner) {
	if broadcast, err := IPV4ToInt(asn.Broadcast); err == nil {
		pipe.ZAdd(ctx, cidrRedisKey, &redis.Z{
			Score:  float64(broadcast),
			Member: asn.CIDR,
		})
	}
	if network, err := IPV4ToInt(asn.Network.IP); err == nil {
		pipe.ZAdd(ctx, cidrRedisKey, &redis.Z{
			Score:  float64(network),
			Member: asn.CIDR,
		})
	}
	pipe.HSet(ctx, asn.GetRedisKey(),
		"org", asn.Org,
		"asn", asn.ASN)
}

// IP Addr Utils

func IPV4ToInt(ipv4 net.IP) (int64, error) {
	if ipv4 == nil {
		return 0, fmt.Errorf("invalid ipv4 address format")
	}
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(ipv4.To4())
	return IPv4Int.Int64(), nil
}

// GeoIP Utils

// The exact limits, as specified by EPSG:900913 / EPSG:3785 / OSGEO:41001 are the following:
// Valid longitudes are from -180 to 180 degrees.
// Valid latitudes are from -85.05112878 to 85.05112878 degrees.

const maxLat = 90
const maxLon = 180

func ValidLat(lat float64) bool {
	return math.Abs(lat) <= maxLat
}

func ValidLon(lon float64) bool {
	return math.Abs(lon) <= maxLon
}

func CoordinateToFloat(coordinateStr string) (float64, error) {
	return strconv.ParseFloat(coordinateStr, 64)
}
