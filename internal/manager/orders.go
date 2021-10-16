package manager

import (
	"github.com/akhmettolegen/onex/pkg/models"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
)

func (m *Manager) GetOrders(page, size int) (response *models.OrdersListResponse, err error) {
	orders, total, err := m.App.DB.GetOrders(page, size)
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

func (m *Manager) CreateOrder(body models.OrderCreateRequest, file *multipart.FileHeader) (response *models.OrderByIDResponse, err error){

	uploadFile, err := m.Upload(file)
	if err != nil {
		return
	}

	filepath := "http://" + m.App.Config.Minio.Host + "/" + m.App.Config.Minio.Bucket + "/" + uploadFile.Data.URL

	orderReq := &models.Order{
		Name:        body.Name,
		Description: body.Description,
		Image:       filepath,
		Status:      models.OrderStatusPending,
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
