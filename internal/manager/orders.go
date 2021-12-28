package manager

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (m *Manager) GetOrders(ui *models.UserInfo, page, size int, me string, statusFilters []models.OrderStatus) (response *models.OrdersListResponse, err error) {
	ifMe := false
	if me == "true" {
		ifMe = true
	}
	orders, total, err := m.App.DB.GetOrders(ui, page, size, ifMe, statusFilters)
	if err != nil {
		return nil, err
	}

	response =  &models.OrdersListResponse{
		Data:  orders,
		Page:  page,
		Size:  size,
		Total: total,
	}
	return
}

func (m *Manager) GetOrderByID(orderID uuid.UUID) (response *models.OrderByIDResponse, err error) {
	order, err := m.App.DB.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	response = &models.OrderByIDResponse{
		Data: order,
	}
	return
}

func (m *Manager) CreateOrder(ui *models.UserInfo, body models.OrderCreateRequest) (response *models.OrderByIDResponse, err error){

	order := models.Order{
		Status:    body.Status,
		TrackCode: body.TrackCode,
		UserID:    ui.UserID,
	}

	err = m.App.DB.CreateOrder(body, &order)
	if err != nil {
		return
	}

	response = &models.OrderByIDResponse{
		Data: order,
	}
	return
}

func (m *Manager) UpdateOrder(req *models.OrderUpdateRequest) (response *models.OrderByIDResponse, err error) {
	order, err := m.App.DB.GetOrderByID(req.ID)
	if err != nil {
		return
	}

	err = m.App.DB.UpdateOrderByID(req)
	if err != nil {
		return
	}

	if req.Product != nil {
		product := &order.Product
		product.PrimeCost = req.Product.PrimeCost
		product.TotalCost = req.Product.TotalCost
		product.InStock = req.Product.InStock
		product.SoldCount = req.Product.SoldCount
		product.Status = req.Product.Status
		product.Image = req.Product.Image
		product.Description = req.Product.Description
		product.Name = req.Product.Name

		err = m.App.DB.UpdateProductByID(product)
		if err != nil {
			return
		}
	}

	order, err = m.App.DB.GetOrderByID(order.ID)
	if err != nil {
		return
	}
	response = &models.OrderByIDResponse{
		Data: order,
	}
	return
}

func (m *Manager) DeleteOrder(orderID uuid.UUID) (err error) {
	return m.App.DB.DeleteOrder(orderID)
}
