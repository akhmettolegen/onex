package gorm

import (
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (g *Gorm) GetUsers(page, size int)(users []models.User, total int, err error) {
	result := g.DB.
		Order("created_at").
		Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}

	offset := (page - 1) * size
	total = len(users)
	err = result.Limit(size).Offset(offset).Find(&users).Error
	return
}

func (g *Gorm) GetUserByID(userId uuid.UUID) (user models.User, err error) {
	err = g.DB.First(&user, userId).Error
	return
}

func (g *Gorm) GetUserByPhone(phone string) (user models.User, err error) {
	err = g.DB.Where("phone = ?", phone).First(&user).Error
	return
}

func (g *Gorm) CreateUser(user *models.User) error {
	return g.DB.Create(&user).Error
}

func (g *Gorm) UpdateUserByID(user *models.User) error {
	return g.DB.Model(&user).Updates(&user).Error
}

func (g *Gorm) DeleteUser(userID uuid.UUID) error {
	return g.DB.Delete(&models.User{}, userID).Error
}
