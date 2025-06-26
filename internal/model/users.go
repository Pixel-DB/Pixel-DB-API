package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=6,max=20"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required,min=2,max=20"`
	LastName  string `json:"lastName" validate:"required,min=2,max=20"`
	Role      string `json:"role" validate:"required,oneof=admin user moderator" gorm:"default:user"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
