package utils

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioClient() (*minio.Client, error) {
	minioClient, err := minio.New("minio-dev:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("root", "iamroot123", ""),
		Secure: false,
	})
	return minioClient, err
}
