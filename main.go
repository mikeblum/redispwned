package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mikeblum/redispwned/api"
	"github.com/mikeblum/redispwned/api/csrf"
	"github.com/mikeblum/redispwned/api/geoip"
	"github.com/mikeblum/redispwned/api/ping"
	"github.com/mikeblum/redispwned/api/report"
	"github.com/mikeblum/redispwned/api/scan"
	config "github.com/mikeblum/redispwned/internal/configs"
)

const envGinMode = "GIN_MODE"
const envCertFile = "CERT_FILE"
const envCertKey = "CERT_KEY"
const modeRelease = "release"

func main() {
	cfg := config.NewConfig()
	log := config.NewLog()
	router := gin.Default()
	api.CORS(router)
	csrf.Routes(router)
	ping.Routes(router)
	report.Routes(router)
	geoip.Routes(router)
	scan.Routes(router)
	var err error
	if mode, ok := os.LookupEnv(envGinMode); ok {
		if strings.EqualFold(mode, modeRelease) {
			err = runRelease(router, cfg)
		}
	} else {
		err = runLocal(router)
	}
	if err != nil {
		log.Fatal("Failed to start router: ", err)
	}
}

func runRelease(router *gin.Engine, cfg *config.AppConfig) error {
	server := http.Server{
		Addr:    "0.0.0.0:443",
		Handler: router,
	}
	return server.ListenAndServeTLS(cfg.GetString(envCertFile), cfg.GetString(envCertKey))
}

func runLocal(router *gin.Engine) error {
	return router.Run() // listen and serve on 0.0.0.0:8080
}
