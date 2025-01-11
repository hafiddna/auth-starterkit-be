package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetConfig(t *testing.T) {
	Config, err := GetConfig()

	t.Logf("Config: %+v", Config)
	assert.NotNil(t, Config)
	assert.Nil(t, err)
}
