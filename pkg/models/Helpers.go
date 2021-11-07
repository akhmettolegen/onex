package models

import (
	"errors"
	"gorm.io/gorm"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// channel validation
	switch u.Channel {
	case ChannelMobile, ChannelWeb:
		return nil
	}
	return errors.New("invalid creative status")
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	// channel validation
	switch o.Status {
	case OrderStatusReady, OrderStatusPending, OrderStatusRecommended:
		return nil
	}
	return errors.New("invalid creative status")
}
