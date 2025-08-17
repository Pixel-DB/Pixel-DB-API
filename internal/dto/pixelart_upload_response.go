package dto

import "time"

type PixelArtUploadDataResponse struct {
	ID                 string
	OwnerUsername      string
	OwnerID            string
	PixelArtURL        string
	PixelArtSize       int64
	CreatedAt          time.Time
	Filename           string
	OldFilename        string
	FileExtension      string
	PixelArtName       string
	PixelArtDesciption string
	PixelArtTags       string
}

type PixelArtUploadResponse struct {
	Status  string
	Message string
	Data    PixelArtUploadDataResponse
}
