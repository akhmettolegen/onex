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
	NetCost int `json:"netCost"`
	Location string `json:"location"`
	DeliveryTime string `json:"deliveryTime"`
	DeliveryCost int `json:"deliveryCost"`
	Warranty string `json:"warranty"`
	Quality string `json:"quality"`
	TotalCost int `json:"totalCost"`
}

type OrdersListResponse struct {
	Data  []Order `json:"data"`
	Page  int     `json:"page"`
	Size  int     `json:"size"`
	Total int     `json:"total"`
}

type OrderCreateRequest struct {
	Name string `form:"name" binding:"required"`
	Description string `form:"description"`
	NetCost int `form:"netCost"`
	Location string `form:"location"`
	DeliveryTime string `form:"deliveryTime"`
	DeliveryCost int `form:"deliveryCost"`
	Warranty string `form:"warranty"`
	Quality string `form:"quality"`
}

type OrderByIDResponse struct {
	Data Order `json:"data"`
}

type OrderUpdateRequest struct {
	ID uuid.UUID `json:"-"`
	Name *string `form:"name" binding:"required"`
	Description *string `form:"description"`
	NetCost *int `form:"netCost"`
	Location *string `form:"location"`
	DeliveryTime *string `form:"deliveryTime"`
	DeliveryCost *int `form:"deliveryCost"`
	Warranty *string `form:"warranty"`
	Quality *string `form:"quality"`
}