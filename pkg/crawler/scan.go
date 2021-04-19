package crawler

import (
	"github.com/zmap/zgrab2"
	zredis "github.com/zmap/zgrab2/modules/redis"
)

const moduleType = "redis"
const scannerName = "RedisScanner"
const defaultRedisPort uint = 6379

type RedisScanner struct {
	zgrab2.Scanner
}

type RedisFlags struct {
	zredis.Flags
}

func (rs *RedisScanner) RedisPort() *uint {
	var redisPort uint = defaultRedisPort
	return &redisPort
}

func NewScanner() *RedisScanner {
	zredis.RegisterModule()
	mod := zgrab2.GetModule(moduleType).(*zredis.Module)
	scan := mod.NewScanner()
	scan.Init(mod.NewFlags().(zgrab2.ScanFlags))
	zgrab2.RegisterScan(moduleType, scan)
	return &RedisScanner{
		scan,
	}
}
