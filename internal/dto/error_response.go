package dto

type ErrorResponse struct {
	Status  string `json:"Status" example:"Error"`
	Message string `json:"Message" example:"Success"`
	Error   string `json:"Error"`
}
