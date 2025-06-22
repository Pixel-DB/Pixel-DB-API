package model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=6,max=20"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=2,max=20"`
	LastName  string `json:"last_name" validate:"required,min=2,max=20"`
	Role      string `json:"role" validate:"required,oneof=admin user moderator" gorm:"default:user"`
}
