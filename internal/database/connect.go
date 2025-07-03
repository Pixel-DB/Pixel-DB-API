package database

import (
	"fmt"
	"strconv"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=db-dev port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(model.Users{}, model.PixelArts{}) // Load the models (tables) to DB
	fmt.Println("Database Migrated")
}
