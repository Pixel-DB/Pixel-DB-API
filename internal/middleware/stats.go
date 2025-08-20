package middleware

import (
	"fmt"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
)

func UpdateRequestCount(count int64) {
	db := database.DB
	db.Model(&model.Stats{}).Where("ID = ?", 1).Update("Count", count)
	fmt.Println(count)
}
