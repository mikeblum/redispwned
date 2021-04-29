package geoip

import (
	"context"
	"net"
	"strconv"

	"github.com/go-redis/redis/v8"
	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/sirupsen/logrus"
)

const plusInf = "+inf"

type Client struct {
	log    *logrus.Entry
	client *redis.Client
}

func NewClient(redisClient *redis.Client) *Client {
	return &Client{
		log:    config.NewLog(),
		client: redisClient,
	}
}

func (geo *Client) LookupIP(ip net.IP) (result []string, err error) {
	ctx := context.Background()
	min, _ := IPV4ToInt(ip)
	zrange := strconv.FormatUint(uint64(min), 10)
	geo.log.Info(zrange)
	cmd := geo.client.ZRangeByScore(ctx, cidrRedisKey, &redis.ZRangeBy{
		Min:    zrange,
		Max:    plusInf,
		Offset: 0,
		Count:  5,
	})
	result, err = cmd.Result()
	if err != nil {
		return nil, err
	}
	geo.log.Debug(result)
	return
}
