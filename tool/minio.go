package tool

import (
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioTool interface {
	Connect() *minio.Client
}

type minioTool struct {
	config config.CfgStruct
}

func NewMinioTool(config config.CfgStruct) MinioTool {
	return &minioTool{
		config: config,
	}
}

func (m *minioTool) Connect() *minio.Client {
	configData := m.config
	endpoint := configData.App.Minio.Host + ":" + configData.App.Minio.Port
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(configData.App.Minio.AccessKey, configData.App.Minio.SecretKey, ""),
		Secure: false, // TODO: Change to true in production
	})
	if err != nil {
		return nil
	}
	return client
}
