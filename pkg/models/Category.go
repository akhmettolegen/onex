package models

import uuid "github.com/satori/go.uuid"

type Category struct {
	Base
	Name 		string `json:"name"`
	Description string `json:"description"`
	Status 		string `json:"status"`
	Slug 		string `json:"slug"`
}

type CategoriesListResponse struct {
	Data  []Category `json:"data"`
	Page  int     	 `json:"page"`
	Size  int     	 `json:"size"`
	Total int     	 `json:"total"`
}

type CategoryCreateRequest struct {
	Name 		string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status 		string `json:"status"`
	Slug 		string `json:"slug"`
}

type CategoryUpdateRequest struct {
	ID 			uuid.UUID `json:"-"`
	Name 		*string   `json:"name"`
	Description *string   `json:"description"`
	Status 		*string   `json:"status"`
	Slug 		*string   `json:"slug"`
}

type CategoryByIDResponse struct {
	Data Category `json:"data"`
}