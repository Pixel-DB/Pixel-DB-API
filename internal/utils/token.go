package utils

import (
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id, email, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  id,
		"email":    email,
		"username": username,
	})

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	fmt.Println("JWT Secret:", config.Config("JWT_SECRET")) // Debugging line to check the secret
	if err != nil {
		return "", err
	}

	return t, nil
}
