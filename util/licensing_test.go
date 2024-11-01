package util

import (
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	licenseConfig = config.NewConfig().GetConfig()

	license = NewLicensing(licenseConfig)
)

func TestLicensing_InitApp(t *testing.T) {
	err := license.InitApp()

	assert.Nil(t, err)
}
