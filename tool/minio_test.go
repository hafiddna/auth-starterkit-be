package tool

import (
	"github.com/hafiddna/auth-starterkit-be/config"
	"testing"
)

var (
	minioConfig = config.NewConfig().GetConfig()

	minioTest = NewMinioTool(minioConfig)
)

func TestMinioTool_Connect(t *testing.T) {
	client := minioTest.Connect()

	if client == nil {
		t.Error("Minio connection failed")
	}
}
