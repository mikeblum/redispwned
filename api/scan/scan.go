package scan

import (
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mikeblum/redispwned/api"
	"github.com/mikeblum/redispwned/internal/geoip"
	"github.com/mikeblum/redispwned/pkg/crawler"
	"github.com/zmap/zgrab2"
)

func Middleware(scan *crawler.RedisScanner) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("scanner", scan)
		c.Next()
	}
}

func Routes(router *gin.Engine) {
	r := router.Group("")
	scanner, _ := crawler.NewScanner()
	r.Use(Middleware(scanner))
	r.POST("/scan/:redisAddr", scanRedis)
}

func scanRedis(c *gin.Context) {
	conn, _ := api.NewRedisConn()
	client := geoip.NewClient(conn)
	redisAddr := c.Param("redisAddr")
	parts := strings.Split(redisAddr, ":")
	ipv4 := net.ParseIP(parts[0])
	port := 6379
	if len(parts) > 1 {
		port, _ = strconv.Atoi(parts[1])
	}
	result, _ := client.LookupIP(ipv4)
	geo, err := client.LookupGeo(result[0])
	if err != nil {
		handleErr(c)
		return
	}
	scanner, _ := c.MustGet("scanner").(*crawler.RedisScanner)
	if err != nil {
		handleErr(c)
		return
	}
	var redisPort uint = uint(port)
	target := zgrab2.ScanTarget{
		IP:   ipv4,
		Port: &redisPort,
	}
	status, info, err := scanner.Scan(target)
	if err != nil {
		handleErr(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"city":         geo["city"],
		"country_code": geo["country_code"],
		"status":       status,
		"info":         info,
	})
}

func handleErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "failed",
	})
}
