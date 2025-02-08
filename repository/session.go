package repository

import (
	"database/sql"
	"github.com/hafiddna/auth-starterkit-be/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(session model.Session, userID string) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(session model.Session, userID string) error {
	if userID == "" {
		session.Created("system")
	} else {
		session.Created(userID)
		session.UserID = sql.NullString{
			String: userID,
			Valid:  true,
		}
	}
	return r.db.Create(&session).Error
}
