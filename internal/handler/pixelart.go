package handler

import (
	"context"
	"encoding/json"
	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/minio/minio-go/v7"
	"github.com/morkid/paginate"
)

// UploadPixelArt godoc
// @Summary      Upload PixelArt
// @Description  Returns a paginated list of pixel art
// @Tags         PixelArt
// @Security BearerAuth
// @Param        pixelart formData  file false "PixelArt-File"
// @consume json
// @Success 200 {object} dto.UploadFileResponse
// @Router /pixelart [post]
func UploadPixelArt(c *fiber.Ctx) error {

	//Get user info from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID := utils.GetUserIDFromToken(token)
	user, err := utils.GetUser(userID)
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Invalid user credentials",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse)
	}

	file, err := c.FormFile("pixelart") //Get file from passed Data
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to get the PixelArt-File",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	metaField := c.FormValue("meta")
	var meta dto.UploadFileRequest
	if err := json.Unmarshal([]byte(metaField), &meta); err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Invalid JSON in 'meta' field",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	fileContent, err := file.Open() //Open file
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to open the PixelArt-File",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	if err := fileContent.Close(); err != nil { //Close File
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to close the PixelArt-File",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	ext := utils.GetExt(file.Filename)
	if ext != "png" {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "FileExtension",
			Error:   "The uploaded file is not a PNG-File",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	// Initialize MinIO client
	minioClient, err := utils.InitMinioClient()
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to initialize storage service",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	newFileName := utils.GenerateFilename(file.Filename, user.Username)

	//Save Data in DB
	pixelArts := &model.PixelArts{
		OwnerID:     user.ID,
		Filename:    newFileName,
		URL:         "https://place-holder-url.com",
		Name:        meta.PixelArtName,
		Description: meta.PixelArtDescription,
	}
	if err := database.DB.Create(pixelArts).Error; err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to create database record",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	//Upload file to MinIO
	_, err = minioClient.PutObject(context.Background(), config.Config("MINIO_BUCKET_NAME"), newFileName, fileContent, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to upload to storage service",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	// Create API response data
	ResponseData := dto.UploadFileResponse{
		Status:  "Success",
		Message: "Uploaded PixelArt-File",
		Data: dto.UploadData{
			ID:                 pixelArts.ID,
			CreatedAt:          pixelArts.CreatedAt,
			OwnerID:            user.ID,
			OwnerUsername:      user.Username,
			Filename:           newFileName,
			OldFilename:        file.Filename,
			FileExtension:      ext,
			PixelArtURL:        "placeholder-url.com", // Placeholder, cooming soon...
			PixelArtSize:       file.Size,
			PixelArtName:       meta.PixelArtName,
			PixelArtDesciption: meta.PixelArtDescription,
		},
	}
	return c.JSON(ResponseData)
}

// GetAllPixelArts godoc
// @Summary      Get pixel art list
// @Description  Returns a paginated list of pixel art
// @Tags         PixelArt
// @Accept       json
// @Produce      json
// @Param        size  query     int  false  "Size of each Page"
// @Param        page  query     int  false  "Page number for pagination"
// @Success      200   {object}   dto.APIResponse
// @Router       /pixelart [get]
func GetAllPixelArts(c *fiber.Ctx) error {
	//All Pixel Arts with Pagination
	var pixelArts []model.PixelArts
	pg := paginate.New()
	data := pg.With(database.DB.Model(&model.PixelArts{})).Request(c.Request()).Response(&pixelArts)

	response := dto.APIResponse{
		Status:  "Success",
		Message: "",
		Data:    data,
	}

	return c.JSON(response)
}

// GetPixelArt godoc
// @Summary Get PixelArt
// @Description  Returns the infos for a specific pixel art by ID
// @Tags PixelArt
// @Param        pixelArtID   path      string  true  "PixelArt ID"
// @Success      200  {file}    file
// @Failure 500 {object} dto.ErrorResponse
// @Router /pixelart/{pixelArtID} [get]
func GetPixelArt(c *fiber.Ctx) error {
	//Single Pixel Art DB entry
	pixelArtID := c.Params("pixelArtID")
	p := new(model.PixelArts)
	db := database.DB
	if err := db.Where(&model.PixelArts{ID: pixelArtID}).First(p).Error; err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Can't find PixelArt",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	response := dto.APIResponse{
		Status:  "Success",
		Message: "",
		Data:    p,
	}
	return c.JSON(response)
}

// GetPixelArtPicture godoc
// @Summary Get PixelArt Picture
// @Description  Returns the image for a specific pixel art by ID
// @Tags PixelArt
// @Param        pixelArtID   path      string  true  "PixelArt ID"
// @Produce png
// @Success      200  {file}    file
// @Failure 500 {object} dto.ErrorResponse
// @Router /pixelart/{pixelArtID}/picture [get]
func GetPixelArtPicture(c *fiber.Ctx) error {

	// get Pixel Art ID from params
	pixelArtID := c.Params("pixelArtID")
	p := new(model.PixelArts)
	db := database.DB
	if err := db.Where(&model.PixelArts{ID: pixelArtID}).First(p).Error; err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Can't find PixelArt",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	// INit Minio
	minioClient, err := utils.InitMinioClient()
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to initialize storage service",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	object, err := minioClient.GetObject(context.Background(), config.Config("MINIO_BUCKET_NAME"), p.Filename, minio.GetObjectOptions{})
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to get pixel art from storage service",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	c.Set("Content-Type", "image/png")

	return c.SendStream(object)
}
