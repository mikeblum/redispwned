package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	assert := assert.New(t)
	cfg := NewConfig()
	assert.NotNil(cfg)
}
