package geoip

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikeblum/redispwned/api"
	"github.com/mikeblum/redispwned/internal/geoip"
)

func Routes(router *gin.Engine) {
	r := router.Group("")
	r.GET("/geoip/:ipv4", geoIP)
}

func geoIP(c *gin.Context) {
	conn, _ := api.NewRedisConn()
	client := geoip.NewClient(conn)
	ipv4 := c.Param("ipv4")
	result, _ := client.LookupIP(net.ParseIP(ipv4))
	geo, err := client.LookupGeo(result[0])
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"city":         geo["city"],
			"country_code": geo["country_code"],
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
		})
	}
}
