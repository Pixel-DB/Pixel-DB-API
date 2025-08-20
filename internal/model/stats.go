package model

import "time"

type Stats struct {
	ID        int64
	Count     int64     `gorm:"not null;default:0"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
