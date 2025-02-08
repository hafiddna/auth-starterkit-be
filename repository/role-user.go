package repository

import "gorm.io/gorm"

type RoleUserRepository interface {
}

type roleUserRepository struct {
	db *gorm.DB
}

func NewRoleUserRepository(db *gorm.DB) RoleUserRepository {
	return &roleUserRepository{db: db}
}
