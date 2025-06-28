package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func UploadPixelArt(c *fiber.Ctx) error {
	file, err := c.FormFile("document")
	if err != nil {
		return err
	}

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to open uploaded file",
			"error":   err.Error(),
		})
	}
	defer fileContent.Close()

	//Setup MinIO client
	minioClient, err := minio.New("minio-dev:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("root", "iamroot123", ""),
		Secure: false,
	})

	_, err = minioClient.PutObject(context.Background(), "test-bucket", "file.txt", fileContent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to upload file to MinIO",
		})
	}

	log.Printf("Successfully uploaded of size %d", file.Size)

	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})

}
