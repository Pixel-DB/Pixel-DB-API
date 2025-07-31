package dto

type ErrorResponse struct {
	Status  string `json:"Status" example:"Error"`
	Message string `json:"Message"`
	Error   string `json:"Error"`
}
