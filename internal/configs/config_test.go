package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	assert := assert.New(t)
	cfg, err := NewConfig()
	assert.Nil(err)
	assert.NotNil(cfg)
}
