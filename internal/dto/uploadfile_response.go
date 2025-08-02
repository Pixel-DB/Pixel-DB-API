package dto

import "time"

type UploadData struct {
	ID            string
	OwnerUsername string
	OwnerID       string
	PixelArtURL   string
	PixelArtSize  int64
	CreatedAt     time.Time
	Filename      string
	OldFilename   string
	FileExtension string
}

type UploadFileResponse struct {
	Status  string
	Message string
	Data    UploadData
}
