package dto

import "time"

type UserData struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	Email     string    `json:"Email"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Username  string    `json:"Username"`
	Role      string    `json:"Role"`
}

type UserResponse struct {
	Status  string   `json:"Status" example:"Success"`
	Message string   `json:"Message" example:"Created User"`
	Data    UserData `json:"Data"`
}
