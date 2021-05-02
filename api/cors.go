package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(router *gin.Engine) {
	allowedOrigins := []string{"http://localhost:1313"}
	apexDomains := []string{"https://haveibeenredised.com", "https://redispwned.app"}
	allowedOrigins = append(allowedOrigins, apexDomains...)
	wwwDomains := []string{"https://www.haveibeenredised.com", "https://www.redispwned.app"}
	allowedOrigins = append(allowedOrigins, wwwDomains...)
	cfDomains := []string{"https://*.redispwned.pages.dev", "https://redispwned.pages.dev"}
	allowedOrigins = append(allowedOrigins, cfDomains...)
	router.Use(cors.New(cors.Config{
		AllowOrigins:  allowedOrigins,
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "Origin", "X-CSRF"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
}
