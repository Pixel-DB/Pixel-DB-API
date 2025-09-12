package handler

import (
	"strings"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/middleware"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/morkid/paginate"
)

func GetAllUsers(c *fiber.Ctx) error {
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
	if !middleware.HasPermission(user.Role, "users.view") {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "You don't have permission for this route",
			Error:   "Forbidden",
		}
		return c.Status(fiber.StatusForbidden).JSON(ErrorResponse)
	}
	var users []dto.UserGetDataResponse

	searchInput := c.Query("search")
	if searchInput != "" {
		pg := paginate.New()
		data := pg.With(database.DB.Model(&model.Users{}).Where("email ILIKE ? OR username ILIKE ? OR last_name ILIKE ? OR first_name ILIKE ?", "%"+strings.ToLower(searchInput)+"%", "%"+strings.ToLower(searchInput)+"%", "%"+strings.ToLower(searchInput)+"%", "%"+strings.ToLower(searchInput)+"%")).Request(c.Request()).Response(&users)
		response := dto.UsersGetAllResponse{
			Status:  "Success",
			Message: "",
			Data:    data,
		}
		return c.JSON(response)
	}

	// Return all Users with Pagination
	pg := paginate.New()
	data := pg.With(database.DB.Model(&model.Users{})).Request(c.Request()).Response(&users)

	response := dto.UsersGetAllResponse{
		Status:  "Success",
		Message: "",
		Data:    data,
	}

	return c.JSON(response)
}
