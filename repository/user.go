package repository

import (
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmailPhoneOrUsername(credential string) (user model.User, err error)
	FindOneById(id string) (user model.User, err error)
	FindByIDWithTokenData(id string) (user model.User, err error)
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
	query := r.db.Where("email = ?", credential).Or("phone = ?", credential).Or("username = ?", credential).
		Preload("Roles", "team_id IS NULL").Preload("Roles.Permissions")
	if config.Config.App.AuthConfig.IsTeamEnabled {
		query = query.Preload("Teams", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN team_user ON team_user.team_id = teams.id").
				Where("team_user.is_active = ?", true)
		}).Preload("Teams.Roles").Preload("Teams.Roles.Permissions").
			Preload("MembersOf", func(db *gorm.DB) *gorm.DB {
				return db.Joins("JOIN team_user ON team_user.team_id = teams.id").
					Where("team_user.is_active = ?", true)
			}).Preload("MembersOf.Roles").Preload("MembersOf.Roles.Permissions")
	}
	err = query.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindOneById(id string) (user model.User, err error) {
	query := r.db.Where("id = ?", id).Preload("Roles", "team_id IS NULL")

	if config.Config.App.AuthConfig.IsTeamEnabled {
		query = query.Preload("Teams").Preload("MembersOf")
	}

	err = query.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByIDWithTokenData(id string) (user model.User, err error) {
	query := r.db.Where("id = ?", id).Preload("Roles", "team_id IS NULL").Preload("Roles.Permissions")
	if config.Config.App.AuthConfig.IsTeamEnabled {
		query = query.Preload("Teams", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN team_user ON team_user.team_id = teams.id").
				Where("team_user.is_active = ?", true)
		}).Preload("Teams.Roles").Preload("Teams.Roles.Permissions").
			Preload("MembersOf", func(db *gorm.DB) *gorm.DB {
				return db.Joins("JOIN team_user ON team_user.team_id = teams.id").
					Where("team_user.is_active = ?", true)
			}).Preload("MembersOf.Roles").Preload("MembersOf.Roles.Permissions")
	}

	err = query.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
