package handler

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/security"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary Login
// @Description Login with your credentials, to get your User Data and your JWT-Token
// @Tags Auth
// @Param        credentials  body  dto.LoginRequest true  "Login Credentials"
// @consume json
// @Success 200 {object} dto.AuthResponse
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	input := new(dto.LoginRequest)

	if err := c.BodyParser(input); err != nil { //Check Request Body
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Invalid Request",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse)
	}

	validate := validator.New() //Validate if Email, ...
	if err := validate.Struct(input); err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Validation Error. Check Request.",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse)
	}

	userModel, err := utils.GetUserEmail(input.Email) //Check if user is in DB
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Error finding user",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse)
	}

	if userModel == nil {
		dummyHash := "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK" //Hash something for Timing Attacks protection
		security.CheckPasswordHash(dummyHash, input.Password)
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Invalid Email or Password",
			Error:   "",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse)
	}

	if !security.CheckPasswordHash(userModel.Password, input.Password) {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Invalid Email or Password",
			Error:   "",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse)
	}

	token, err := utils.GenerateToken(userModel.ID, userModel.Email, userModel.Username)
	if err != nil {
		ErrorResponse := dto.ErrorResponse{
			Status:  "Error",
			Message: "Invalid Email or Password",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse)
	}

	AuthResponse := dto.AuthResponse{
		ID:       userModel.ID,
		Email:    userModel.Email,
		Username: userModel.Username,
		Role:     userModel.Role,
		Token:    token,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Logged in", "data": AuthResponse})
}
