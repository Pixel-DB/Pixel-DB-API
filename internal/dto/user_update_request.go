package dto

type UpdateUserRequest struct {
	Username  string `gorm:"uniqueIndex" json:"username" validate:"required,min=3,max=20,alphanum"`
	Password  string `json:"password" validate:"required,min=6,max=70"`
	Email     string `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required,min=2,max=20,alpha"`
	LastName  string `json:"lastName" validate:"required,min=2,max=20,alpha"`
}
