package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikeblum/redispwned/api"
)

func Routes(router *gin.Engine) {
	ping := router.Group("")
	ping.GET("/health", healthCheck)
	ping.GET("/ping", healthCheck)
}

func healthCheck(c *gin.Context) {
	conn, ctx := api.NewRedisConn()
	if pong, err := conn.Ping(ctx).Result(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": pong,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"staus": "failed",
		})
	}
}
