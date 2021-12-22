package models

import (
	"errors"
	"gorm.io/gorm"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// user channel validation
	switch u.Channel {
	case ChannelMobile, ChannelWeb:
		return nil
	}
	return errors.New("invalid user channel")
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	// order status validation
	switch o.Status {
	case OrderStatusAnalyzed, OrderStatusNotAnalyzed:
		return nil
	}
	return errors.New("invalid order status")
}

func (o *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// order status validation
	switch o.Status {
	case ProductStatusActive, ProductStatusInactive:
		return nil
	}
	return errors.New("invalid product status")
}