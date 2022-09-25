package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role"`
}