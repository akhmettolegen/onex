package models

import uuid "github.com/satori/go.uuid"

type ProductStatus string
const (
	ProductStatusActive ProductStatus = "ACTIVE"
	ProductStatusInactive ProductStatus = "INACTIVE"
)

type Product struct {
	Base
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Status ProductStatus `json:"status"`
	SoldCount int `json:"soldCount"`
	InStock bool `json:"inStock" gorm:"default:false"`
	PrimeCost float64 `json:"primeCost"`
	TotalCost float64 `json:"totalCost"`
}

type ProductsListResponse struct {
	Data  []Product `json:"data"`
	Page  int     `json:"page"`
	Size  int     `json:"size"`
	Total int     `json:"total"`
}

type ProductCreateRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Status ProductStatus `json:"status"`
	SoldCount int `json:"soldCount"`
	InStock bool `json:"inStock" gorm:"default:false"`
	PrimeCost float64 `json:"primeCost"`
	TotalCost float64 `json:"totalCost"`
}

type ProductByIDResponse struct {
	Data Product `json:"data"`
}

type ProductUpdateRequest struct {
	ID uuid.UUID `json:"-"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Image *string `json:"image"`
	Status *ProductStatus `json:"status"`
	SoldCount *int `json:"soldCount"`
	InStock *bool `json:"inStock" gorm:"default:false"`
	PrimeCost *float64 `json:"primeCost"`
	TotalCost *float64 `json:"totalCost"`
}