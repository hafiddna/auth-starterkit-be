package repository

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(session entity.Session) interface{}
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(session entity.Session) interface{} {
	return r.db.Create(&session).Error
}
