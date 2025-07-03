package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIDFromToken(t *jwt.Token) string {
	claims := t.Claims.(jwt.MapClaims)
	id := string(claims["user_id"].(string))

	return id
}
