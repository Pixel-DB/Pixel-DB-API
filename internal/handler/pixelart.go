package handler

import (
	"context"
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/minio/minio-go/v7"
)

func UploadPixelArt(c *fiber.Ctx) error {

	//Get user info from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID := utils.GetUserIDFromToken(token)
	user, err := utils.GetUser(userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user credentials",
			"error":   err.Error(),
		})
	}

	file, err := c.FormFile("pixelart") //Get file from passed Data
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to upload the file",
		})
	}

	fileContent, err := file.Open() //Open file
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to open uploaded file",
		})
	}
	defer fileContent.Close()

	//Generate a new filename
	newFilename := utils.GenerateFilename(file.Filename, user.Username)

	// Initialize MinIO client
	minioClient, err := utils.InitMinioClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to initialize storage service",
			"error":   err.Error(),
		})
	}

	//Upload file to MinIO
	_, err = minioClient.PutObject(context.Background(), config.Config("MINIO_BUCKET_NAME"), newFilename, fileContent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to upload file to S3-Service",
		})
	}

	//Save Data in DB
	pixelArts := &model.PixelArts{
		OwnerID: user.ID,
	}

	if err := database.DB.Create(pixelArts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create database record",
			"error":   err.Error(),
		})
	}

	// Create API response data
	data := dto.UploadFileResponse{
		ID:            pixelArts.ID,
		CreatedAt:     pixelArts.CreatedAt,
		OwnerID:       user.ID,
		OwnerUsername: user.Username,
		Filename:      newFilename,
		OldFilename:   file.Filename,
		PixelArtURL:   "placeholder-url.com", // Placeholder, cooming soon...
		PixelArtSize:  file.Size,
	}
	fmt.Println(user.ID, user.Email)

	return c.JSON(fiber.Map{"status": "success", "message": "Uploaded the File", "data": data})

}
