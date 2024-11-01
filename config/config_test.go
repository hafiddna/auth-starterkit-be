package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetConfig(t *testing.T) {
	configTest := NewConfig()
	cfg := configTest.GetConfig()

	t.Logf("Config: %+v", cfg)
	assert.NotNil(t, cfg)
}
