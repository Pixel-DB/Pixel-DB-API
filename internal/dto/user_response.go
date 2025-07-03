package dto

import "time"

type UserResponse struct {
	ID        string
	CreatedAt time.Time
	Username  string
	Email     string
	FirstName string
	LastName  string
	Role      string
}
