package database

import (
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectToMinio() (client *minio.Client, err error) {
	endpoint := config.Config.App.Minio.Host + ":" + config.Config.App.Minio.Port
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Config.App.Minio.AccessKey, config.Config.App.Minio.SecretKey, ""),
		Secure: false, // TODO: Change to true in production
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
