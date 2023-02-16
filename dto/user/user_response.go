package userdto

import (
	"dumbsound/models"
)

type UserResponse struct {
	ID       int         `json:"id"`
	RoleID   int         `json:"roleid"`
	Role     models.Role `json:"role"`
	Fullname string      `json:"fullname" gorm:"type: varchar(255)"`
	Email    string      `json:"email" gorm:"type: varchar(255)"`
	Password string      `json:"password" gorm:"type: varchar(255)"`
	Gender   string      `json:"gender" gorm:"type: varchar(255)"`
	Phone    string      `json:"phone" gorm:"type: varchar(255)"`
	Address  string      `json:"address" gorm:"type: varchar(255)"`
}
