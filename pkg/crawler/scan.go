package crawler

import (
	"time"

	"github.com/zmap/zgrab2"
	zredis "github.com/zmap/zgrab2/modules/redis"
)

const moduleType = "redis"
const scannerName = "redis"
const defaultRedisPort uint = 6379

type RedisScanner struct {
	*zredis.Scanner
}

func (rs *RedisScanner) RedisPort() *uint {
	var redisPort uint = defaultRedisPort
	return &redisPort
}

// https://github.com/zmap/zgrab2/blob/master/modules/redis.go
func init() {
	zredis.RegisterModule()
}

func NewScanner() *RedisScanner {
	mod := zgrab2.GetModule(moduleType).(*zredis.Module)
	scan := mod.NewScanner().(*zredis.Scanner)
	// configure flags
	flags := mod.NewFlags().(*zredis.Flags)
	flags.Port = defaultRedisPort
	flags.Timeout = time.Duration(10) * time.Second // seconds
	flags.Name = scannerName
	scan.Init(flags)
	zgrab2.RegisterScan(moduleType, scan)
	return &RedisScanner{
		scan,
	}
}
