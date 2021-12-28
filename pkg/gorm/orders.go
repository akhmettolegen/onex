package gorm

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func (g *Gorm) GetOrders(ui *models.UserInfo, page, size int, me bool, statusFilters []models.OrderStatus) (orders []models.Order, total int, err error) {
	result := g.DB.
		Order("created_at")

	if me {
		result = result.Where("user_id = ?", ui.UserID)
	}

	result = result.
		Where("status in ?", statusFilters).
		Preload("Product").
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
	err = g.DB.Preload("Product").First(&order, orderId).Error
	return
}

func (g *Gorm) CreateOrder(req models.OrderCreateRequest, order *models.Order) error {
	err := g.DB.Transaction(func(tx *gorm.DB) error {
		product := models.Product{
			Description: req.Product.Description,
			Name: 		 req.Product.Name,
			Image:       req.Product.Image,
			Status:      req.Product.Status,
			SoldCount:   req.Product.SoldCount,
			InStock:     req.Product.InStock,
			PrimeCost:   req.Product.PrimeCost,
			TotalCost:   req.Product.TotalCost,
		}
		err := tx.Clauses(dbresolver.Write).Create(&product).Error
		if err != nil {
			return err
		}
		order.Product = product
		order.ProductID = product.ID

		err = tx.Clauses(dbresolver.Write).Create(&order).Error
		return err
	})
	return err
}

func (g *Gorm) UpdateOrderByID(req *models.OrderUpdateRequest) error {
	order := models.Order{}
	order.ID = req.ID

	if req.Status != nil {
		order.Status = *req.Status
	}
	if req.TrackCode != nil {
		order.TrackCode = *req.TrackCode
	}
	return g.DB.Model(&order).Updates(&order).Error
}

func (g *Gorm) DeleteOrder(orderID uuid.UUID) error {
	return g.DB.Delete(&models.Order{}, orderID).Error
}