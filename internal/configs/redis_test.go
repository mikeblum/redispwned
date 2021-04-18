package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisDefaults(t *testing.T) {
	assert := assert.New(t)
	expectedAddr := "127.0.0.1:6379"
	assert.Equal(expectedAddr, defaultRedisAddr)
	expectedDB := 0
	assert.Equal(expectedDB, defaultRedisDB)
}
