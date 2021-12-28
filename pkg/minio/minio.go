package minioClient

import (
	"fmt"
	"github.com/akhmettolegen/texert/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Get - MinIOClient initializer
func Get(cfg *config.Config) (*minio.Client, error) {
	endpoint := fmt.Sprintf("%v", cfg.Minio.Host)
	accessKeyID := cfg.Minio.AccessKey
	secretAccessKey := cfg.Minio.SecretKey
	useSSL := cfg.Minio.UseSSL

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println("init minio client", err)
		panic(1)
		return nil, err
	}

	fmt.Println("client:", client) // minioClient is now set up

	return client, nil
}

