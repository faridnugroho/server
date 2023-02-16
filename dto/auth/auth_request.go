package authdto

import (
	"dumbsound/models"
	"time"
)

type RegisterRequest struct {
	RoleID    int         `json:"roleid" form:"roleid" gorm:"type: varchar(255)"`
	Role      models.Role `json:"role"`
	Fullname  string      `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Email     string      `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password  string      `json:"password" form:"password" gorm:"type: varchar(255)"`
	Gender    string      `json:"gender" form:"gender" gorm:"type: varchar(255)"`
	Phone     string      `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Address   string      `json:"address" form:"address" gorm:"type: varchar(255)"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255)"  validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
