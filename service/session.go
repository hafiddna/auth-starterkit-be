package service

import (
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type SessionService interface {
	Create(session model.Session, userID string) error
}

type sessionService struct {
	sessionRepository repository.SessionRepository
}

func NewSessionService(sessionRepository repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: sessionRepository,
	}
}

func (s *sessionService) Create(session model.Session, userID string) error {
	return s.sessionRepository.Create(session, userID)
}
