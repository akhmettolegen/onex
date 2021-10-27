package manager

import (
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (m *Manager) GetOrders(ui *models.UserInfo, page, size int, me string) (response *models.OrdersListResponse, err error) {
	ifMe := false
	if me == "true" {
		ifMe = true
	}
	orders, total, err := m.App.DB.GetOrders(ui, page, size, ifMe)
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

	response =  &models.OrderByIDResponse{
		Data: order,
	}
	return
}

func (m *Manager) CreateOrder(ui *models.UserInfo, body models.OrderCreateRequest) (response *models.OrderByIDResponse, err error){

	orderReq := &models.Order{
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
		Status:      models.OrderStatusPending,
		UserID:      ui.UserID,
		NetCost:	 body.NetCost,
		Location: 	 body.Location,
		DeliveryTime: body.DeliveryTime,
		DeliveryCost: body.DeliveryCost,
		Warranty:     body.Warranty,
		Quality: 	  body.Quality,
		TotalCost:    body.NetCost + body.DeliveryCost,
	}

	err = m.App.DB.CreateOrder(orderReq)
	if err != nil {
		return
	}

	order, err := m.App.DB.GetOrderByID(orderReq.ID)
	if err != nil {
		return
	}

	response = &models.OrderByIDResponse{
		Data: order,
	}
	return
}

func (m *Manager) UpdateOrder(req models.OrderUpdateRequest) (response *models.OrderByIDResponse, err error) {
	order, err := m.App.DB.GetOrderByID(req.ID)
	if err != nil {
		return
	}

	if req.Description != nil {
		order.Description = *req.Description
	}
	if req.Name != nil {
		order.Name = *req.Name
	}
	if req.NetCost != nil {
		order.NetCost = *req.NetCost
	}
	if req.Location != nil {
		order.Location = *req.Location
	}
	if req.DeliveryTime != nil {
		order.DeliveryTime = *req.DeliveryTime
	}
	if req.DeliveryCost != nil {
		order.DeliveryCost = *req.DeliveryCost
	}
	if req.Quality != nil {
		order.Quality = *req.Quality
	}
	if req.Warranty != nil {
		order.Warranty = *req.Warranty
	}

	err = m.App.DB.UpdateOrderByID(&order)
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
