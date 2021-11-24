package gorm

import (
	"fmt"
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (g *Gorm) GetOrders(ui *models.UserInfo, page, size int, me bool, statusFilters []models.OrderStatus) (orders []models.Order, total int, err error) {
	result := g.DB.
		Order("created_at")

	if me {
		result = result.Where("user_id = ?", ui.UserID)
	}

	result = result.
		Where("status in ?", statusFilters).
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

func (g *Gorm) GetOrders2() (orders []models.Order, total int, err error) {
	db, err := g.DB.DB()
	rows, err := db.Query("SELECT status, description FROM orders LIMIT $1", 3)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var status string
		var description string
		err := rows.Scan(&status, &description)
		if err != nil {
			panic(err)
		}
		fmt.Println(status, description)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
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