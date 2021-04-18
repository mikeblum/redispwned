package shodan

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisMeta struct {
	Hash      int        `json:"hash"`
	IP        int        `json:"ip"`
	IPStr     string     `json:"ip_str"`
	ISP       string     `json:"isp"`
	Location  Location   `json:"location"`
	Org       string     `json:"org"`
	Port      int        `json:"port"`
	Product   string     `json:"product"`
	RawData   string     `json:"data"`
	Redis     RedisData  `json:"redis"`
	Shodan    ShodanMeta `json:"_shodan"`
	Transport string     `json:"transport"`
	Version   string     `json:"version"`
}

func (rm *RedisMeta) GetRedisKey() string {
	return fmt.Sprintf("shodan:%s", rm.Shodan.ID)
}

func (rm *RedisMeta) GetGeo() string {
	return fmt.Sprintf("%f,%f", rm.Location.Latitude, rm.Location.Longitude)
}

func (rm *RedisMeta) ToHSet(ctx context.Context, pipe redis.Pipeliner) {
	pipe.HSet(ctx, rm.GetRedisKey(),
		"hash", rm.Hash,
		"ip", rm.IP,
		"ip_addr", rm.IPStr,
		"port", rm.Port,
		"isp", rm.ISP,
		"geo", rm.GetGeo(),
		"city", rm.Location.City,
		"country", rm.Location.CountryName,
		"country_code", rm.Location.CountryCode,
		"redis_version", rm.Version,
		"region_code", rm.Location.RegionCode,
		"raw_data", rm.RawData,
		"uptime_days", rm.Redis.Server.UptimeDays,
		"uptime_seconds", rm.Redis.Server.UptimeSeconds)
}

type ShodanMeta struct {
	ID string `json:"id"`
}

type RedisData struct {
	Stats   RedisStats   `json:"stats"`
	Cluster RedisCluster `json:"cluster"`
	Server  RedisServer  `json:"server"`
}

type RedisStats struct {
	TotalConnectionsProcessed int `json:"total_connections_received"`
	RejectedConnections       int `json:"rejected_connections"`
}

type RedisCluster struct {
	Enabled bool `json:"cluster_enabled"`
}

func (rc *RedisCluster) UnmarshalJSON(data []byte) error {
	raw := make(map[string]interface{})
	err := json.Unmarshal(data, &raw)
	if err != nil {
		rc.Enabled = false
		return nil
	}
	if clusterEnabled, ok := raw["cluster_enabled"]; ok {
		rc.Enabled = clusterEnabled != 0
	}
	return err
}

type RedisServer struct {
	UptimeSeconds int `json:"uptime_in_seconds"`
	UptimeDays    int `json:"uptime_in_days"`
}

type Location struct {
	City        string  `json:"city"`
	RegionCode  string  `json:"region_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CountryName string  `json:"country_name"`
	CountryCode string  `json:"country_code"`
}
