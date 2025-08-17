package dto

import "time"

type UserGetDataResponse struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	Email     string    `json:"Email"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Username  string    `json:"Username"`
	Role      string    `json:"Role"`
}

type UserGetResponse struct {
	Status  string              `json:"Status" example:"Success"`
	Message string              `json:"Message" example:"Get User"`
	Data    UserGetDataResponse `json:"Data"`
}
