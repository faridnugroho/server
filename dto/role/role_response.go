package roledto

import (
	"time"
)

type RoleResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
