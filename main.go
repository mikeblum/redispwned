package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mikeblum/redispwned/api"
	"github.com/mikeblum/redispwned/api/ping"
	"github.com/mikeblum/redispwned/api/report"
	config "github.com/mikeblum/redispwned/internal/configs"
)

func main() {
	log := config.NewLog()
	router := gin.Default()
	api.CORS(router)
	ping.Routes(router)
	report.Routes(router)
	err := router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("Failed to start router: ", err)
	}
}
