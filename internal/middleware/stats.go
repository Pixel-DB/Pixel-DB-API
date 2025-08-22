package middleware

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
)

func UpdateRequestCount(count int64) {
	db := database.DB
	var stats model.Stats

	db.Where("id = ?", 1).First(&stats)

	stats.RequestCount = stats.RequestCount + count
	db.Save(&stats)
}
