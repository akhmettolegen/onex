package application

import (
	"context"
	"fmt"
	"github.com/akhmettolegen/onex/pkg/config"
	"github.com/akhmettolegen/onex/pkg/gorm"
	minioClient "github.com/akhmettolegen/onex/pkg/minio"
	"github.com/minio/minio-go/v7"
)

// Application model
type Application struct {
	Config *config.Config
	DB *gorm.Gorm
	MinIOClient *minio.Client
}

// Get - Application initializer
func Get() (*Application, error) {
	config := config.Get()
	dbManager := gorm.Get(config)
	// Init MinIOClient
	client, err := minioClient.Get(config)
	if err != nil {
		fmt.Println("minio client init err", err)
	}

	buckets, err := client.ListBuckets(context.Background())
	if err != nil {
		fmt.Println("list buckets error", err)
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	if config.DB.AutoMigrate {
		dbManager.AutoMigrate()
	}
	return &Application{
		Config: config,
		DB: &dbManager,
		MinIOClient: client,
	}, nil
}
