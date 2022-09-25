package authdto

type LoginResponse struct {
	ID        int    `json:"id"`
	FullName  string `gorm:"type: varchar(255)" json:"fullName"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Gender    string `gorm:"type: varchar(255)" json:"gender"`
	Phone     string `gorm:"type: varchar(255)" json:"phone"`
	Address   string `gorm:"type: varchar(255)" json:"address"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Subscribe bool   `json:"subscribe" form:"subscribe"`
	Role      string `gorm:"type: varchar(100)" json:"role"`
}

type CheckAuthResponse struct {
	ID        int    `json:"id"`
	FullName  string `gorm:"type: varchar(255)" json:"fullName"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Gender    string `gorm:"type: varchar(255)" json:"gender"`
	Phone     string `gorm:"type: varchar(255)" json:"phone"`
	Address   string `gorm:"type: varchar(255)" json:"address"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Subscribe bool   `json:"subscribe" form:"subscribe"`
	Role      string `gorm:"type: varchar(100)" json:"role"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Role     string `gorm:"type: varchar(100)" json:"role"`
}