package utils

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
)

func DeleteUser(userID string) error {
	db := database.DB

	if err := db.Delete(&model.Users{}, "id = ?", userID).Error; err != nil {
		return err
	}

	return nil
}
