package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID   string
	Email    string
	Username string
}

func GetUserFromToken(t *jwt.Token) UserClaims {
	claims := t.Claims.(jwt.MapClaims)
	id := string(claims["user_id"].(string))
	email := string(claims["email"].(string))
	username := string(claims["username"].(string))

	return UserClaims{UserID: id, Email: email, Username: username}
}
