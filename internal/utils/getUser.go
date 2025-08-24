package utils

import (
	"errors"
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"gorm.io/gorm"
)

func GetUser(i string) (*model.Users, error) { //Get user by ID
	u := new(model.Users)
	db := database.DB
	err := db.First(u, "id = ?", i).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User with ID %s not found", i)
		}
		return nil, err
	}
	return u, nil
}
