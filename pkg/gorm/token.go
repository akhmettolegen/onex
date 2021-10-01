package gorm

import "github.com/akhmettolegen/onex/pkg/models"

func (g *Gorm) CreateToken(token *models.AccessToken) error {
	return g.DB.Create(&token).Error
}
