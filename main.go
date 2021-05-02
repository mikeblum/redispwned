package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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
	configureLogging(router)
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
	return router.Run() // default to 0.0.0.0:8080
}

func configureLogging(router *gin.Engine) {
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
}
