package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindRoles() ([]models.Role, error)
	GetRole(ID int) (models.Role, error)
	CreateRole(role models.Role) (models.Role, error)
	UpdateRole(role models.Role) (models.Role, error)
	DeleteRole(role models.Role) (models.Role, error)
}

func RepositoryRole(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindRoles() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error

	return roles, err
}

func (r *repository) GetRole(ID int) (models.Role, error) {
	var role models.Role
	err := r.db.First(&role, ID).Error

	return role, err
}

func (r *repository) CreateRole(role models.Role) (models.Role, error) {
	err := r.db.Create(&role).Error

	return role, err
}

func (r *repository) UpdateRole(role models.Role) (models.Role, error) {
	err := r.db.Save(&role).Error

	return role, err
}

func (r *repository) DeleteRole(role models.Role) (models.Role, error) {
	err := r.db.Delete(&role).Error

	return role, err
}
