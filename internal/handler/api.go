package handler

import (
	"context"
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func UploadPixelArt(c *fiber.Ctx) error {
	//Setup MinIO client
	minioClient, err := minio.New("minio-dev:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("root", "iamroot123", ""),
		Secure: false,
	})

	file, err := c.FormFile("document") //Get file from passed Data
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

	newFilename := utils.GenerateFilename(file.Filename) //Generate a new filename

	//Upload file to MinIO
	_, err = minioClient.PutObject(context.Background(), "test-bucket", newFilename, fileContent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to upload file to S3-Service",
		})
	}

	token := c.Locals("user").(*jwt.Token)    //Load Token
	userID := utils.GetUserIDFromToken(token) //Get UserID from Token
	user, err := utils.GetUser(userID)        //Get User from DB

	db := database.DB //Get DB instance
	PixelArts := new(model.PixelArts)

	PixelArts.OwnerID = user.ID
	if err := db.Create(PixelArts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "error": err.Error()})
	}

	data := dto.UploadFileResponse{
		ID:            PixelArts.ID,
		CreatedAt:     PixelArts.CreatedAt,
		OwnerID:       user.ID,
		OwnerUsername: user.Username,
		Filename:      newFilename,
		OldFilename:   file.Filename,
		PixelArtURL:   "placeholder-url.com",
		PixelArtSize:  file.Size,
	}
	fmt.Println(user.ID, user.Email)

	return c.JSON(fiber.Map{"status": "success", "message": "Uploaded the File", "data": data})

}
