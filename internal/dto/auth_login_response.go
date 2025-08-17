package dto

import "time"

type AuthLoginDataResponse struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	Email     string    `json:"Email"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Username  string    `json:"Username"`
	Role      string    `json:"Role"`
}

type AuthLoginResponse struct {
	Status  string                `json:"Status" example:"Success"`
	Message string                `json:"Message" example:"Logged in"`
	Token   string                `json:"Token"`
	Data    AuthLoginDataResponse `json:"Data"`
}
