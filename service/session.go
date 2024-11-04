package service

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type SessionService interface {
	Create(session entity.Session) interface{}
}

type sessionService struct {
	sessionRepository repository.SessionRepository
}

func NewSessionService(sessionRepository repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: sessionRepository,
	}
}

func (s *sessionService) Create(session entity.Session) interface{} {
	return s.sessionRepository.Create(session)
}
