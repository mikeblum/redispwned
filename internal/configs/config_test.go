package config

import (
	"testing"

	asserts "github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	assert := asserts.New(t)
	cfg := NewConfig()
	assert.NotNil(cfg)
}
