package gorm

import (
	"fmt"
	"github.com/akhmettolegen/onex/pkg/config"
	"github.com/akhmettolegen/onex/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBManager struct {
	DB *gorm.DB
}

func Get(config *config.Config) DBManager {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v TimeZone=Asia/Almaty", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Pass, config.DB.Mode)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}

	return DBManager{
		DB: db,
	}
}

func (g *DBManager) AutoMigrate() {
	fmt.Println("migrating ...")
	g.DB.AutoMigrate(
		models.Base{},
	)
}
