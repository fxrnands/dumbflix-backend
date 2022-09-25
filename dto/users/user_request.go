package usersdto

type CreateUserRequest struct {
	FullName  string `json:"fullName" form:"fullName" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Subscribe bool   `json:"subscribe" form:"subscribe"`
	Role      string `json:"role"`
}

type UpdateUserRequest struct {
	FullName  string `json:"fullName" form:"fullName"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Gender    string `json:"gender" form:"gender"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	Subscribe bool   `json:"subscribe" form:"subscribe"`
	Role      string `json:"role"`
}