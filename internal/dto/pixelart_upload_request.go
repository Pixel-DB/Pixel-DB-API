package dto

type PixelArtUploadRequest struct {
	PixelArtTitle       string `json:"PixelArtName" validate:"required,min=3,max=70"`
	PixelArtDescription string `json:"PixelArtDescription" validate:"required,min=6,max=200"`
}
