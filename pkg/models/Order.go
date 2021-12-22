package models

import uuid "github.com/satori/go.uuid"

type OrderStatus string
const (
	OrderStatusNotAnalyzed OrderStatus = "NOT_ANALYZED"
	OrderStatusAnalyzed OrderStatus = "ANALYZED"
)

type Order struct {
	Base
	Product Product `json:"product"`
	ProductID uuid.UUID `json:"productId"`
	Status OrderStatus `json:"status" gorm:"default:NOT_ANALYZED"`
	TrackCode string `json:"trackCode"`
	UserID uuid.UUID `json:"userId"`
}

type OrdersListResponse struct {
	Data  []Order `json:"data"`
	Page  int     `json:"page"`
	Size  int     `json:"size"`
	Total int     `json:"total"`
}

type OrderCreateRequest struct {
	Product Product `json:"product"`
	ProductID uuid.UUID `json:"productId"`
	Status OrderStatus `json:"status" gorm:"default:NOT_ANALYZED"`
	TrackCode string `json:"trackCode"`
}

type OrderByIDResponse struct {
	Data Order `json:"data"`
}

type OrderUpdateRequest struct {
	ID uuid.UUID `json:"-"`
	Product *Product `json:"product"`
	ProductID *uuid.UUID `json:"productId"`
	Status *OrderStatus `json:"status" gorm:"default:NOT_ANALYZED"`
	TrackCode *string `json:"trackCode"`
}