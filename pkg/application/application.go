package application

import (
	"github.com/akhmettolegen/onex/pkg/config"
	"github.com/akhmettolegen/onex/pkg/gorm"
)

// Application model
type Application struct {
	Config *config.Config
	DB *gorm.DBManager
}

// Get - Application initializer
func Get() (*Application, error) {
	config := config.Get()
	dbManager := gorm.Get(config)

	if config.DB.AutoMigrate {
		dbManager.AutoMigrate()
	}
	return &Application{
		Config: config,
		DB: &dbManager,
	}, nil
}
