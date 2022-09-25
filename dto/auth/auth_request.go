package authdto

type RegisterRequest struct {
	FullName string `gorm:"type: varchar(255)" json:"fullName" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender   string `gorm:"type: varchar(100)" json:"gender" validate:"required"`
	Phone    string `gorm:"type: varchar(100)" json:"phone" validate:"required"`
	Address  string `gorm:"type: text" json:"address" validate:"required"`
	Role     string `gorm:"type: varchar(100)" json:"role"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}