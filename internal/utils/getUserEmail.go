package utils

import (
	"errors"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"gorm.io/gorm"
)

func GetUserEmail(e string) (*model.Users, error) {
	u := new(model.Users)
	db := database.DB
	if err := db.Where(&model.Users{Email: e}).First(u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
