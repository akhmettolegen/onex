package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type AccessToken struct {
	Token uuid.UUID `json:"token" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID uuid.UUID `json:"userId" gorm:"uniqueIndex"`
	CreatedAt *time.Time `json:"createdAt" gorm:"default:now()"`
	TTL int `json:"ttl"`
}

type AccessTokenCreateRequest struct {
	Token uuid.UUID `json:"token"`
	UserID uuid.UUID `json:"userId"`
	TTL int `json:"ttl"`
}


type SignUpResponse struct {
	Token uuid.UUID `json:"token"`
	UserID uuid.UUID `json:"userId"`
	TTL int `json:"ttl"`
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	Token uuid.UUID `json:"token"`
	UserID uuid.UUID `json:"userId"`
	TTL int `json:"ttl"`
}

type SignInRequest struct {
	Phone string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckTokenResponse struct {
	Token     uuid.UUID  `json:"token"`
	UserID    uuid.UUID  `json:"user_id"`
}