package model

import "time"

type Stats struct {
	ID           int64
	RequestCount int64
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
