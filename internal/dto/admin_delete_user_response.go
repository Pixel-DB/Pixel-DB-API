package dto

type AdminDeleteUserDataResponse struct {
	ID string `json:"ID"`
}

type AdminDeleteUserResponse struct {
	Status  string                      `json:"Status" example:"Success"`
	Message string                      `json:"Message" example:"Updated User"`
	Data    AdminDeleteUserDataResponse `json:"Data"`
}
