package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/mikeblum/redispwned/internal/configs"
)

func main() {
	log := config.NewLog()
	routes := gin.Default()
	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := routes.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("Failed to start router: ", err)
	}
}
