package gorm

import (
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (g *Gorm) GetOrders(ui *models.UserInfo, page, size int, me bool) (orders []models.Order, total int, err error) {
	result := g.DB.
		Order("created_at")

	if me {
		result = result.Where("user_id = ?", ui.UserID)
	}

	result = result.
		Find(&orders)

	if result.Error != nil {
		err = result.Error
		return
	}

	offset := (page - 1) * size
	total = len(orders)
	err = result.Limit(size).Offset(offset).Find(&orders).Error
	return
}

func (g *Gorm) GetOrderByID(orderId uuid.UUID) (order models.Order, err error) {
	err = g.DB.First(&order, orderId).Error
	return
}

func (g *Gorm) CreateOrder(order *models.Order) error {
	return g.DB.Create(&order).Error
}

func (g *Gorm) UpdateOrderByID(order *models.Order) error {
	return g.DB.Model(&order).Updates(&order).Error
}

func (g *Gorm) DeleteOrder(orderID uuid.UUID) error {
	return g.DB.Delete(&models.Order{}, orderID).Error
}