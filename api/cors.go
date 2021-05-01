package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:1313", "https://haveibeenredised.com", "https://redispwned.app"},
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "Origin", "X-CSRF"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
}
