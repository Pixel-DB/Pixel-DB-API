package middleware

import (
	"fmt"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
)

func UpdateRequestCount(count int64) {
	db := database.DB
	c := new(model.Stats)

	db.Where(&model.Stats{ID: 1}).First(c)
	oldCount := c.Count
	fmt.Print("Old")
	fmt.Println(oldCount)

	db.Model(&model.Stats{}).Where("ID = ?", 1).Update("Count", oldCount+count)
	fmt.Println(count)
}
