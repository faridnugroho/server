package roledto

import (
	"time"
)

type CreateRoleRequest struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateRoleRequest struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
