package csrf

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/mikeblum/redispwned/api"
)

const csrfRedisKey = "csrf:"

func Routes(router *gin.Engine) {
	csrf := router.Group("")
	csrf.GET("/csrf", csrfToken)
}

func CheckCSRFToken(ctx context.Context, conn *redis.Client, csrfToken string) error {
	_, err := conn.Get(ctx, csrfRedisKey+csrfToken).Result()
	return err
}

func csrfToken(c *gin.Context) {
	conn, ctx := api.NewRedisConn()
	token := uuid.Must(uuid.NewRandom())
	if _, err := conn.Set(ctx, csrfRedisKey+token.String(), token, expireIn()).Result(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
		})
	}
}

func expireIn() time.Duration {
	return time.Hour
}
