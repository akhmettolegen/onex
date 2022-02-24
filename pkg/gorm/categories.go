package gorm

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (g *Gorm) GetCategories(ui *models.UserInfo, page, size int) (categories []models.Category, total int, err error) {
	result := g.DB.
		Order("created_at").
		Find(&categories)

	if result.Error != nil {
		err = result.Error
		return
	}

	offset := (page - 1) * size
	total = len(categories)
	err = result.Limit(size).Offset(offset).Find(&categories).Error
	return
}

func (g *Gorm) GetCategoryByID(categoryId uuid.UUID) (category models.Category, err error) {
	err = g.DB.First(&category, categoryId).Error
	return
}

func (g *Gorm) CreateCategory(category *models.Category) error {
	return g.DB.Create(&category).Error
}

func (g *Gorm) UpdateCategoryByID(category *models.Category) error {
	return g.DB.Model(&category).Updates(&category).Error
}

func (g *Gorm) DeleteCategory(categoryID uuid.UUID) error {
	return g.DB.Delete(&models.Category{}, categoryID).Error
}