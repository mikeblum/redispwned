package shodan

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	IPStr    string   `json:"ip_str"`
	ISP      string   `json:"isp"`
	Location Location `json:"location"`
	ASN      string   `json:"asn"`
	Org      string   `json:"org"`
	Version  string   `json:"version"`
}

func (rm *Redis) GetRedisKey() string {
	return "shodan:" + rm.IPStr
}

func (rm *Redis) GetGeo() string {
	return fmt.Sprintf("%f,%f", rm.Location.Latitude, rm.Location.Longitude)
}

func (rm *Redis) ToHSet(ctx context.Context, pipe redis.Pipeliner) {
	pipe.HSet(ctx, rm.GetRedisKey(),
		"isp", rm.ISP,
		"asn", rm.ASN,
		"geo", rm.GetGeo(),
		"city", rm.Location.City,
		"country", rm.Location.CountryName,
		"country_code", rm.Location.CountryCode,
		"version", rm.Version)
}

type Location struct {
	City        string  `json:"city"`
	RegionCode  string  `json:"region_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CountryName string  `json:"country_name"`
	CountryCode string  `json:"country_code"`
}
