package model

import "time"

type Stats struct {
	ID        int64
	Count     int64
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
