package service

import (
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type SessionService interface {
	CreateOrUpdate(session model.Session) error
}

type sessionService struct {
	sessionRepository repository.SessionRepository
}

func NewSessionService(sessionRepository repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: sessionRepository,
	}
}

func (s *sessionService) CreateOrUpdate(session model.Session) error {
	err := s.sessionRepository.FindOneByUserID(session.UserID.String)
	if err != nil {
		return s.sessionRepository.Create(session)
	} else {
		return s.sessionRepository.Update(session)
	}
}
