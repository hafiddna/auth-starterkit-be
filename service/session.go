package service

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type SessionService interface {
	Create(session entity.Session, userID string) error
}

type sessionService struct {
	sessionRepository repository.SessionRepository
}

func NewSessionService(sessionRepository repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: sessionRepository,
	}
}

func (s *sessionService) Create(session entity.Session, userID string) error {
	return s.sessionRepository.Create(session, "")
}
