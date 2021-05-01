package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mikeblum/redispwned/api"
	"github.com/mikeblum/redispwned/api/csrf"
	"github.com/mikeblum/redispwned/api/geoip"
	"github.com/mikeblum/redispwned/api/ping"
	"github.com/mikeblum/redispwned/api/report"
	"github.com/mikeblum/redispwned/api/scan"
	config "github.com/mikeblum/redispwned/internal/configs"
)

func main() {
	log := config.NewLog()
	router := gin.Default()
	api.CORS(router)
	csrf.Routes(router)
	ping.Routes(router)
	report.Routes(router)
	geoip.Routes(router)
	scan.Routes(router)
	err := router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("Failed to start router: ", err)
	}
}
