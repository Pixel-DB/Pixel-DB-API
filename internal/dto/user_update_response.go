package dto

type UserUpdateResponse struct {
	Status  string   `json:"Status" example:"Success"`
	Message string   `json:"Message" example:"Updated User"`
	Data    UserData `json:"Data"`
}

type UserUpdateData struct {
	ID        string `json:"ID"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Role      string `json:"Role"`
}
