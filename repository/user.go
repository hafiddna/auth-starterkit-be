package repository

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmailPhoneOrUsername(credential string) (user entity.User, err error)
	FindOneById(id string) (user entity.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmailPhoneOrUsername(credential string) (user entity.User, err error) {
	err = r.db.Where("email = ?", credential).Or("phone = ?", credential).Or("username = ?", credential).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindOneById(id string) (user entity.User, err error) {
	err = r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
