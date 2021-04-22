package crawler

import (
	"time"

	"github.com/zmap/zgrab2"
	zredis "github.com/zmap/zgrab2/modules/redis"
)

const moduleType = "redis"
const scannerName = "redis"
const defaultRedisPort uint = 6379
const defaultTimeoutSeconds int = 10 // seconds

type RedisScanner struct {
	*zredis.Scanner
}

func (rs *RedisScanner) RedisPort() *uint {
	var redisPort uint = defaultRedisPort
	return &redisPort
}

// https://github.com/zmap/zgrab2/blob/master/modules/redis.go
func NewScanner() (*RedisScanner, error) {
	zredis.RegisterModule()
	mod := zgrab2.GetModule(moduleType).(*zredis.Module)
	scan := mod.NewScanner().(*zredis.Scanner)
	// configure flags
	flags := mod.NewFlags().(*zredis.Flags)
	flags.Port = defaultRedisPort
	flags.Timeout = time.Duration(defaultTimeoutSeconds) * time.Second // seconds
	flags.Name = scannerName
	err := scan.Init(flags)
	if err != nil {
		return nil, err
	}
	zgrab2.RegisterScan(moduleType, scan)
	return &RedisScanner{
		scan,
	}, nil
}
