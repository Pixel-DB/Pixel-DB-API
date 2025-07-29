package utils

import (
	"errors"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"gorm.io/gorm"
)

func GetUser(i string) (*model.Users, error) { //Get user by ID
	u := new(model.Users)
	db := database.DB
	if err := db.Where(&model.Users{ID: i}).First(u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
