package repository

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmailPhoneOrUsername(credential string) entity.User
	FindOneById(id string) entity.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmailPhoneOrUsername(credential string) entity.User {
	var user entity.User
	r.db.Where("email = ? OR phone = ? OR username = ?", credential, credential, credential).Preload("Roles").Preload("Roles.Permissions").First(&user)
	return user
}

func (r *userRepository) FindOneById(id string) entity.User {
	var user entity.User
	r.db.Where("id = ?", id).First(&user)
	return user
}
