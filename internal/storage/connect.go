package storage

import (
	"log"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/gofiber/storage/minio"
)

func InitMinioClient() error {

	storage := minio.New(minio.Config{
		Bucket:   "pixel-arts",
		Endpoint: "minio-dev:9000",
		Credentials: minio.Credentials{
			AccessKeyID:     config.Config("MINIO_USER"),
			SecretAccessKey: config.Config("MINIO_PASSWORD"),
		},
	})

	if err := storage.CheckBucket(); err != nil {
		if err := storage.CreateBucket(); err != nil {
			log.Fatalf("failed to create bucket: %v", err)
		}
	}
	return nil
}
