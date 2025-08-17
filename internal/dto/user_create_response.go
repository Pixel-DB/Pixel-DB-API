package dto

import "time"

type UserCreateDataResponse struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	Email     string    `json:"Email"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Username  string    `json:"Username"`
	Role      string    `json:"Role"`
}

type UserCreateResponse struct {
	Status  string                 `json:"Status" example:"Success"`
	Message string                 `json:"Message" example:"Created User"`
	Data    UserCreateDataResponse `json:"Data"`
}
