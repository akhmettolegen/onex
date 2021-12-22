package gorm

import (
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (g *Gorm) GetProducts(ui *models.UserInfo, page, size int, statusFilters []models.ProductStatus) (products []models.Product, total int, err error) {
	result := g.DB.
		Order("created_at")

	result = result.
		Where("status in ?", statusFilters).
		Find(&products)

	if result.Error != nil {
		err = result.Error
		return
	}

	offset := (page - 1) * size
	total = len(products)
	err = result.Limit(size).Offset(offset).Find(&products).Error
	return
}

func (g *Gorm) GetProductByID(productId uuid.UUID) (product models.Product, err error) {
	err = g.DB.First(&product, productId).Error
	return
}

func (g *Gorm) CreateProduct(product *models.Product) error {
	return g.DB.Create(&product).Error
}

func (g *Gorm) UpdateProductByID(product *models.Product) error {
	return g.DB.Model(&product).Updates(&product).Error
}

func (g *Gorm) DeleteProduct(productID uuid.UUID) error {
	return g.DB.Delete(&models.Product{}, productID).Error
}