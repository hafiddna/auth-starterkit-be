package repository

import (
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmailPhoneOrUsername(credential string) (user model.User, err error)
	FindOneById(id string) (user model.User, err error)
}

type userRepository struct {
	db    *gorm.DB
	minio *minio.Client
}

func NewUserRepository(db *gorm.DB, minio *minio.Client) UserRepository {
	return &userRepository{
		db:    db,
		minio: minio,
	}
}

func (r *userRepository) FindByEmailPhoneOrUsername(credential string) (user model.User, err error) {
	err = r.db.Where("email = ?", credential).Or("phone = ?", credential).Or("username = ?", credential).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindOneById(id string) (user model.User, err error) {
	err = r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
