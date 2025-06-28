package handler

import (
	"context"
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/gofiber/fiber/v2"
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

	file, err := c.FormFile("document")
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to upload the file",
		})
	}

	fileContent, err := file.Open()
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to open uploaded file",
		})
	}
	defer fileContent.Close()

	_, err = minioClient.PutObject(context.Background(), "test-bucket", utils.GenerateFilename(file.Filename), fileContent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to upload file to S3-Service",
		})
	}

	//token := c.Locals("user").(*jwt.Token)
	//u, err := utils.GetUser(i.Email)

	return c.JSON(fiber.Map{"status": "success", "message": "Uploaded the File", "data": nil})

}
