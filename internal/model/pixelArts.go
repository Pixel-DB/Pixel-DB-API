package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PixelArts struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	URL       string `validate:"required,url"`
	OwnerID   string `validate:"required" gorm:"type:uuid;"`
	Filename  string `validate:"required"`
}

func (u *PixelArts) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
