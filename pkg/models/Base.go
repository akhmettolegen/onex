package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)
type Base struct {
	ID 		   uuid.UUID 	  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt  *time.Time 	  `json:"createdAt" gorm:"default:now()"`
	UpdatedAt  *time.Time 	  `json:"updatedAt" gorm:"default:now()"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" swaggertype:"string"`
}
