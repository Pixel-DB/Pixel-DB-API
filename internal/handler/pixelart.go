package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/go-playground/validator"
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
// @Success 200 {object} dto.PixelArtUploadResponse
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
	meta := new(dto.PixelArtUploadRequest)
	if err := json.Unmarshal([]byte(metaField), &meta); err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Invalid JSON in 'meta' field",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	validate := validator.New()
	if err := validate.Struct(meta); err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Validation Error. Check Request.",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse)
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
	defer fileContent.Close() //Close File, when function Ends

	buf, err := io.ReadAll(fileContent)
	if err != nil {
		return err
	}

	res, err := utils.CheckResolution(bytes.NewReader(buf))
	fmt.Println(res)

	ext := utils.GetExt(file.Filename) //Get file extension
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
		Title:       meta.PixelArtTitle,
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
	_, err = minioClient.PutObject(context.Background(), config.Config("MINIO_BUCKET_NAME"), newFileName, bytes.NewReader(buf), int64(len(buf)), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Failed to upload to storage service",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	// Create API response data
	PixelArtUploadResponse := dto.PixelArtUploadResponse{
		Status:  "Success",
		Message: "Uploaded PixelArt-File",
		Data: dto.PixelArtUploadDataResponse{
			ID:                 pixelArts.ID,
			CreatedAt:          pixelArts.CreatedAt,
			OwnerID:            user.ID,
			OwnerUsername:      user.Username,
			Filename:           newFileName,
			OldFilename:        file.Filename,
			FileExtension:      ext,
			PixelArtURL:        "placeholder-url.com", // Placeholder, cooming soon...
			PixelArtSize:       file.Size,
			PixelArtName:       meta.PixelArtTitle,
			PixelArtDesciption: meta.PixelArtDescription,
		},
	}
	return c.JSON(PixelArtUploadResponse)
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

	response := dto.PixelArtGetAllResponse{
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

	response := dto.PixelArtGetAllResponse{
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
