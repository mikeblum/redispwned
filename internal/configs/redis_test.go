package config

import (
	"testing"

	asserts "github.com/stretchr/testify/assert"
)

func TestRedisDefaults(t *testing.T) {
	assert := asserts.New(t)
	expectedAddr := "127.0.0.1:6379"
	assert.Equal(expectedAddr, defaultRedisAddr)
	expectedDB := 0
	assert.Equal(expectedDB, DefaultRedisDB)
}

func TestRedisDBConfigurable(t *testing.T) {
	assert := asserts.New(t)
	expectedDB := 1
	client := newRedisClient(expectedDB)
	assert.Equal(expectedDB, client.Options().DB)
}
