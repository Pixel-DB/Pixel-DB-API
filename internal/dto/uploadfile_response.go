package dto

import "time"

type UploadFileResponse struct {
	ID            string
	OwnerUsername string
	OwnerID       string
	PixelArtURL   string
	PixelArtSize  int64
	CreatedAt     time.Time
	Filename      string
	OldFilename   string
}
