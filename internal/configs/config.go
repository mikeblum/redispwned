package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const envLogLevel = "LOG_LEVEL"

var defaultLogLevel = logrus.InfoLevel.String()

const envLogFormat = "LOG_FORMAT"

var jsonLogFormat = "JSON"

type AppConfig struct {
	*viper.Viper
}

func NewConfig() *AppConfig {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("..")
	config.AddConfigPath("../..")
	config.AddConfigPath("../../..")
	config.AddConfigPath("/home/redis/.secrets/redispwned")
	config.AutomaticEnv()
	_ = config.ReadInConfig()
	return &AppConfig{
		config,
	}
}

func NewLog() *logrus.Entry {
	var log = logrus.New()
	var cfg = NewConfig()
	if val, ok := HasEnv(cfg.GetString(envLogFormat)); ok {
		switch strings.ToUpper(val) {
		case jsonLogFormat:
			log.SetFormatter(&logrus.JSONFormatter{
				DisableHTMLEscape: true,
			})
		default:
			log.SetFormatter(&logrus.TextFormatter{})
		}
	}
	log.SetOutput(os.Stdout)
	if val, ok := HasEnv(cfg.GetString(envLogLevel)); ok {
		level, err := logrus.ParseLevel(val)
		if err != nil {
			level, _ = logrus.ParseLevel(defaultLogLevel)
		}
		log.SetLevel(level)
	}
	return log.WithFields(logrus.Fields{})
}

func HasEnv(env string) (string, bool) {
	return env, len(env) > 0
}

func GetEnv(envVar, defaultValue string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	}
	return defaultValue
}
