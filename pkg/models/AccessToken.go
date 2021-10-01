package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type AccessToken struct {
	Token uuid.UUID `json:"token"`
	UserID uuid.UUID `json:"userId"`
	CreatedAt *time.Time `json:"createdAt"`
	TTL int `json:"ttl"`
}
