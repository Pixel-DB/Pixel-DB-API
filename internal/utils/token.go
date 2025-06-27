package utils

import (
	"time"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id, email, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  id,
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}
