package dto

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role"`
}
