package handler

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/minio/minio-go/v7"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
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

	// Initialize MinIO client
	minioClient, err := utils.InitMinioClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to initialize storage service",
			"error":   err.Error(),
		})
	}

	newFileName := utils.GenerateFilename(file.Filename, user.Username)

	//Save Data in DB
	pixelArts := &model.PixelArts{
		OwnerID:  user.ID,
		Filename: newFileName,
	}

	if err := database.DB.Create(pixelArts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create database record",
			"error":   err.Error(),
		})
	}

	//Upload file to MinIO
	_, err = minioClient.PutObject(context.Background(), config.Config("MINIO_BUCKET_NAME"), newFileName, fileContent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to upload file to S3-Service",
		})
	}

	// Create API response data
	data := dto.UploadFileResponse{
		ID:            pixelArts.ID,
		CreatedAt:     pixelArts.CreatedAt,
		OwnerID:       user.ID,
		OwnerUsername: user.Username,
		Filename:      newFileName,
		OldFilename:   file.Filename,
		PixelArtURL:   "placeholder-url.com", // Placeholder, cooming soon...
		PixelArtSize:  file.Size,
	}
	fmt.Println(user.ID, user.Email)

	return c.JSON(fiber.Map{"status": "success", "message": "Uploaded the File", "data": data})

}

func GetPixelArt(c *fiber.Ctx) error {

	//All Pixel Arts with Pagination
	if c.Params("pixelArtID") == "" {
		var pixelArts []model.PixelArts
		pg := paginate.New()

		res := pg.With(database.DB.Model(&model.PixelArts{})).Request(c.Request()).Response(&pixelArts)

		return c.JSON(fiber.Map{
			"data": res,
		})
	}

	//Single Pixel Art DB entry
	pixelArtID := c.Params("pixelArtID")
	p := new(model.PixelArts)
	db := database.DB
	if err := db.Where(&model.PixelArts{ID: pixelArtID}).First(p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   p,
	})
}

func GetPixelArtPicture(c *fiber.Ctx) error {

	// get Pixel Art ID from params
	pixelArtID := c.Params("pixelArtID")
	p := new(model.PixelArts)
	db := database.DB
	if err := db.Where(&model.PixelArts{ID: pixelArtID}).First(p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}

	// INit Minio
	minioClient, err := utils.InitMinioClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to initialize storage service",
			"error":   err.Error(),
		})
	}

	//
	object, err := minioClient.GetObject(context.Background(), config.Config("MINIO_BUCKET_NAME"), p.Filename, minio.GetObjectOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to get pixel art from storage service"})
	}

	ext := strings.ToLower(strings.TrimPrefix(path.Ext(p.Filename), ".")) //To Lower, Cut ".", get Ext
	c.Set("Content-Type", "image/"+ext)

	return c.SendStream(object)

}
