package repository

import (
	"database/sql"
	"github.com/hafiddna/auth-starterkit-be/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	FindOneByUserID(userID string) error
	Create(session model.Session) error
	Update(session model.Session) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) FindOneByUserID(userID string) error {
	return r.db.Where("user_id = ?", userID).First(&model.Session{}).Error
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
	return r.db.Create(&session).Error
}

func (r *sessionRepository) Update(session model.Session) error {
	return r.db.Save(&session).Error
}
