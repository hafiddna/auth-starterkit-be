package repository

import (
	"database/sql"
	"github.com/hafiddna/auth-starterkit-be/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	FindOneByAppID(appID string) (session model.Session, err error)
	Create(session model.Session) error
	Update(session model.Session) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) FindOneByAppID(appID string) (session model.Session, err error) {
	err = model.WithoutTrashed(r.db).
		Where("app_id = ?", appID).
		First(&session).Error
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
