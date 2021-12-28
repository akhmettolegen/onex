package manager

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (m *Manager) GetProducts(ui *models.UserInfo, page, size int, statusFilters []models.ProductStatus) (response *models.ProductsListResponse, err error) {
	products, total, err := m.App.DB.GetProducts(ui, page, size, statusFilters)
	if err != nil {
		return nil, err
	}

	response =  &models.ProductsListResponse{
		Data:  products,
		Page:  page,
		Size:  size,
		Total: total,
	}
	return
}

func (m *Manager) GetProductByID(productID uuid.UUID) (response *models.ProductByIDResponse, err error) {
	product, err := m.App.DB.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	response =  &models.ProductByIDResponse{
		Data: product,
	}
	return
}

func (m *Manager) CreateProduct(ui *models.UserInfo, body models.ProductCreateRequest) (response *models.ProductByIDResponse, err error){

	productReq := &models.Product{
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
		Status:      body.Status,
		SoldCount:   body.SoldCount,
		InStock:     body.InStock,
		PrimeCost:   body.PrimeCost,
		TotalCost:   body.TotalCost,
	}

	err = m.App.DB.CreateProduct(productReq)
	if err != nil {
		return
	}

	product, err := m.App.DB.GetProductByID(productReq.ID)
	if err != nil {
		return
	}

	response = &models.ProductByIDResponse{
		Data: product,
	}
	return
}

func (m *Manager) UpdateProduct(req models.ProductUpdateRequest) (response *models.ProductByIDResponse, err error) {
	product, err := m.App.DB.GetProductByID(req.ID)
	if err != nil {
		return
	}

	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.PrimeCost != nil {
		product.PrimeCost = *req.PrimeCost
	}
	if req.TotalCost != nil {
		product.TotalCost = *req.TotalCost
	}
	if req.InStock != nil {
		product.InStock = *req.InStock
	}
	if req.SoldCount != nil {
		product.SoldCount = *req.SoldCount
	}
	if req.Image != nil {
		product.Image = *req.Image
	}
	if req.Status != nil {
		product.Status = *req.Status
	}

	err = m.App.DB.UpdateProductByID(&product)
	if err != nil {
		return
	}
	response = &models.ProductByIDResponse{
		Data: product,
	}
	return
}

func (m *Manager) DeleteProduct(productID uuid.UUID) (err error) {
	return m.App.DB.DeleteProduct(productID)
}
