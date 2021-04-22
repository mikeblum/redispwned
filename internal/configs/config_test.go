package config

import (
	"testing"

	asserts "github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	assert := asserts.New(t)
	cfg, err := NewConfig()
	assert.Nil(err)
	assert.NotNil(cfg)
}
