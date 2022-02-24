package manager

import (
	"github.com/akhmettolegen/texert/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func (m *Manager) GetCategories(ui *models.UserInfo, page, size int) (response *models.CategoriesListResponse, err error) {
	categories, total, err := m.App.DB.GetCategories(ui, page, size)
	if err != nil {
		return nil, err
	}

	response =  &models.CategoriesListResponse{
		Data:  categories,
		Page:  page,
		Size:  size,
		Total: total,
	}
	return
}

func (m *Manager) GetCategoryByID(categoryID uuid.UUID) (response *models.CategoryByIDResponse, err error) {
	category, err := m.App.DB.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	response =  &models.CategoryByIDResponse{
		Data: category,
	}
	return
}

func (m *Manager) CreateCategory(ui *models.UserInfo, body models.CategoryCreateRequest) (response *models.CategoryByIDResponse, err error){

	categoryReq := &models.Category{
		Name:        body.Name,
		Description: body.Description,
		Slug:        body.Slug,
		Status:      body.Status,
	}

	err = m.App.DB.CreateCategory(categoryReq)
	if err != nil {
		return
	}

	category, err := m.App.DB.GetCategoryByID(categoryReq.ID)
	if err != nil {
		return
	}

	response = &models.CategoryByIDResponse{
		Data: category,
	}
	return
}

func (m *Manager) UpdateCategory(req models.CategoryUpdateRequest) (response *models.CategoryByIDResponse, err error) {
	category, err := m.App.DB.GetCategoryByID(req.ID)
	if err != nil {
		return
	}

	if req.Description != nil {
		category.Description = *req.Description
	}
	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Status != nil {
		category.Status = *req.Status
	}
	if req.Slug != nil {
		category.Slug = *req.Slug
	}

	err = m.App.DB.UpdateCategoryByID(&category)
	if err != nil {
		return
	}
	response = &models.CategoryByIDResponse{
		Data: category,
	}
	return
}

func (m *Manager) DeleteCategory(categoryID uuid.UUID) (err error) {
	return m.App.DB.DeleteCategory(categoryID)
}
