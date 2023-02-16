package authdto

import "dumbsound/models"

type RegisterResponse struct {
	Email string `json:"email" gorm:"type: varchar(255)" `
	Token string `json:"token" gorm:"type: varchar(255)" `
}

type LoginResponse struct {
	Email string `json:"email" gorm:"type: varchar(255)" `
	Token string `json:"token" gorm:"type: varchar(255)" `
}

type CheckAuthResponse struct {
	ID    int         `gorm:"type: int" json:"id"`
	Email string      `json:"email" gorm:"type: varchar(255)" `
	Role  models.Role `json:"role"`
}
