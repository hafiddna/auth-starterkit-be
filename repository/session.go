package repository

import (
	"database/sql"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	FindOneByAppID(appID string, preloadTokenData bool) (session model.Session, err error)
	Create(session model.Session) error
	Update(session model.Session) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) FindOneByAppID(appID string, preloadTokenData bool) (session model.Session, err error) {
	query := model.WithoutTrashed(r.db).Where("app_id = ?", appID)

	if preloadTokenData {
		query = query.Preload("User").Preload("User.Roles", "team_id IS NULL").Preload("User.Roles.Permissions")
		if config.Config.App.AuthConfig.IsTeamEnabled {
			query = query.Preload("User.Teams", func(db *gorm.DB) *gorm.DB {
				return db.Joins("JOIN team_user ON team_user.team_id = teams.id").
					Where("team_user.is_active = ?", true)
			}).Preload("User.Teams.Roles").Preload("User.Teams.Roles.Permissions").
				Preload("User.MembersOf", func(db *gorm.DB) *gorm.DB {
					return db.Joins("JOIN team_user ON team_user.team_id = teams.id").
						Where("team_user.is_active = ?", true)
				}).Preload("User.MembersOf.Roles").Preload("User.MembersOf.Roles.Permissions")
		}
	}

	err = query.First(&session).Error
	if err != nil {
		return session, err
	}

	return session, nil
}

func (r *sessionRepository) Create(session model.Session) error {
	if session.UserID.String == "" {
		session.Created("system")
	} else {
		session.Created(session.UserID.String)
		session.UserID = sql.NullString{
			String: session.UserID.String,
			Valid:  true,
		}
	}
	session.Created(session.UserID.String)
	return r.db.Create(&session).Error
}

func (r *sessionRepository) Update(session model.Session) error {
	session.Updated(session.UserID.String)
	return r.db.Save(&session).Error
}
