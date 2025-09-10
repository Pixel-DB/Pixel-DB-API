package handler

import (
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/security"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/morkid/paginate"
)

// CreateUser godoc
// @Summary Create User
// @Description Creates a new user account by accepting user details such as username, email, password, firstname and lastname.
// @Tags User
// @Param        credentials  body  dto.UserCreateRequest true  "User Credentials"
// @consume json
// @Success 200 {object} dto.UserCreateResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user [post]
func CreateUser(c *fiber.Ctx) error {
	u := new(model.Users)
	r := dto.UserCreateRequest{}
	db := database.DB

	if err := c.BodyParser(&r); err != nil {
		return c.JSON(fiber.Map{"status": "Error"})
	}

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Validation Error. Check Request.",
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

	UserCreateResponse := dto.UserCreateResponse{
		Status:  "Success",
		Message: "Created User",
		Data: dto.UserCreateDataResponse{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Role:      u.Role,
		},
	}

	return c.JSON(UserCreateResponse)
}

// GetUser godoc
// @Summary Get User
// @Description Get the user Data, when passing your JWT-Token in the Heeader
// @Tags User
// @Security BearerAuth
// @Success 200 {object} dto.UserGetResponse
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

	UserGetResponse := dto.UserGetResponse{
		Status:  "success",
		Message: "Get User",
		Data: dto.UserGetDataResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
		},
	}

	return c.JSON(UserGetResponse)
}

// UpdateUser godoc
// @Summary Update User
// @Description Update the User Data
// @Tags User
// @Security BearerAuth
// @Success 200 {object} dto.UserUpdateResponse
// @Router /user [patch]
func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	r := dto.UserUpdateRequest{}
	if err := c.BodyParser(&r); err != nil {
		return c.JSON(fiber.Map{"status": "Error"})
	}

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Validation Error. Check Request.",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse)
	}

	token := c.Locals("user").(*jwt.Token)
	userID := utils.GetUserIDFromToken(token)

	updates := map[string]interface{}{}
	if r.FirstName != "" {
		updates["first_name"] = r.FirstName
	}
	if r.LastName != "" {
		updates["last_name"] = r.LastName
	}
	if r.Email != "" {
		updates["email"] = r.Email
	}
	if r.Username != "" {
		updates["username"] = r.Username
	}
	if r.Password != "" {
		hashedPassword, err := security.HashPassword(r.Password)
		if err != nil {
			ErrorResponse := dto.ErrorResponse{
				Status:  "Error",
				Message: "Couldn't Hash new Password",
				Error:   err.Error(),
			}
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
		}
		updates["password"] = hashedPassword
	}

	fmt.Println(updates)
	db.Model(&model.Users{}).Where("id = ?", userID).Updates(updates)

	user, err := utils.GetUser(userID)
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Can't fetch User",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse)
	}

	UserUpdateResponse := dto.UserUpdateResponse{
		Status:  "Success",
		Message: "Updated User",
		Data: dto.UserUpdateDataResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
		},
	}

	return c.JSON(UserUpdateResponse)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []dto.UserGetDataResponse

	pg := paginate.New()
	data := pg.With(database.DB.Model(&model.Users{})).Request(c.Request()).Response(&users)

	response := dto.UsersGetAllResponse{
		Status:  "Success",
		Message: "",
		Data:    data,
	}

	return c.JSON(response)
}
