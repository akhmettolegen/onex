package gorm

import (
	"fmt"
	"github.com/akhmettolegen/onex/pkg/config"
	"github.com/akhmettolegen/onex/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	DB *gorm.DB
}

func Get(config *config.Config) Gorm {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v TimeZone=Asia/Almaty", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Pass, config.DB.Mode)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}

	return Gorm{
		DB: db,
	}
}

func (g *Gorm) AutoMigrate() {
	fmt.Println("migrating ...")
	g.DB.AutoMigrate(
		models.Base{},
		models.User{},
		models.Company{},
		models.AccessToken{},
		models.Product{},
		models.Order{},
	)
}
