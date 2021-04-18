package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	log := NewLog()
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("..")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn("Fatal error config file: ", err)
	}
	return config
}

func NewLog() *logrus.Entry {
	var log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	return log.WithFields(logrus.Fields{})
}

func GetEnv(envVar string, defaultValue string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	}
	return defaultValue
}
