package handler

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/security"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// CreateUser godoc
// @Summary Create User
// @Description Creates a new user account by accepting user details such as username, email, password, firstname and lastname.
// @Tags User
// @Param        credentials  body  dto.UserRequest true  "User Credentials"
// @consume json
// @Success 200 {object} dto.UserResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user [post]
func CreateUser(c *fiber.Ctx) error {
	u := new(model.Users)
	r := dto.UserRequest{}
	db := database.DB

	if err := c.BodyParser(&r); err != nil {
		return c.JSON(fiber.Map{"status": "Error"})
	}

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Validation Error. Check Request.-",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse)
	}

	u.Username = r.Username
	u.Email = r.Email
	u.FirstName = r.FirstName
	u.LastName = r.LastName
	u.Role = "user"

	hashedPassword, err := security.HashPassword(r.Password)
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Couldn't Hash Password",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}
	u.Password = hashedPassword

	if err := db.Create(u).Error; err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Couldn't create User",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusConflict).JSON(ErrorResponse)
	}

	UserResponse := dto.UserResponse{
		Status:  "Success",
		Message: "Created User",
		Data: dto.UserData{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Role:      u.Role,
		},
	}

	return c.JSON(UserResponse)

}

// GetUser godoc
// @Summary Get User
// @Description Get the user Data, when passing your JWT-Token in the Heeader
// @Tags User
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Router /user [get]
func GetUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	userID := utils.GetUserIDFromToken(token)
	user, err := utils.GetUser(userID)
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Can't fetch User",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse)
	}

	response := dto.UserResponse{
		Status:  "success",
		Message: "Get User",
		Data: dto.UserData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
		},
	}

	return c.JSON(response)
}

// UpdateUser godoc
// @Summary Update User
// @Description Update the User Data
// @Tags User
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Router /user [patch]
func UpdateUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	userID := utils.GetUserIDFromToken(token)
	user, err := utils.GetUser(userID)
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Can't fetch User",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse)
	}

	return c.JSON(fiber.Map{"Token": token, "UserID": userID, "User": user})
}
