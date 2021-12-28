package gorm

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func (g *Gorm) CreateToken(token *models.AccessToken) error {
	return g.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"token", "ttl", "created_at"}),
	}).Create(&token).Error
}

func (g *Gorm) GetAccessToken(token uuid.UUID) (*models.AccessToken, error) {
	tkn := &models.AccessToken{}
	err := g.DB.Where("token = ?", token).First(tkn).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, models.AccessTokenNotFoundError
	} else if err != nil {
		return nil, err
	}
	if tkn.CreatedAt.Add(time.Second * time.Duration(tkn.TTL)).Before(time.Now()) {
		err = g.DeleteAccessToken(token)
		if err != nil {
			return nil, err
		}
		return nil, models.AccessTokenExpiredError
	}
	return tkn, nil
}

func (g *Gorm) DeleteAccessToken(token uuid.UUID) error {
	res := g.DB.Where("token = ?", token).Delete(&models.AccessToken{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected <= 0 {
		return models.AccessTokenNotFoundError
	}
	return nil
}