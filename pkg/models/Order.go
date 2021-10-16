package models

import uuid "github.com/satori/go.uuid"

type OrderStatus string
const (
	OrderStatusReady OrderStatus = "READY"
	OrderStatusPending OrderStatus = "PENDING"
)

type Order struct {
	Base
	Name string `json:"name"`
	Description string `json:"description"`
	Status OrderStatus `json:"status" gorm:"default:PENDING"`
	Image string `json:"image"`
	UserID uuid.UUID `json:"userId"`
}

type OrdersListResponse struct {
	Data  []Order `json:"data"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
}

type OrderCreateRequest struct {
	Name string `form:"name" binding:"required"`
	Description string `form:"description"`
}

type OrderByIDResponse struct {
	Data Order `json:"data"`
}

type OrderUpdateRequest struct {
	ID uuid.UUID `json:"-"`
	Name *string `form:"name" binding:"required"`
	Description *string `form:"description"`
}