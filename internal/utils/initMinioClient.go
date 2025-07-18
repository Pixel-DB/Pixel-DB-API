package utils

import (
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioClient() (*minio.Client, error) {
	minioClient, err := minio.New("minio-dev:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(config.Config("MINIO_USER"), config.Config("MINIO_PASSWORD"), ""),
		Secure: false,
	})
	fmt.Println("Initializing MinIO client...")
	return minioClient, err
}
