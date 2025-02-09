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
	sessionData, err := s.sessionRepository.FindOneByUserIDAndUserAgent(session.UserID.String, session.UserAgent.String)
	if err != nil || session.UserAgent.String != sessionData.UserAgent.String {
		return s.sessionRepository.Create(session)
	} else {
		return s.sessionRepository.Update(session)
	}
}
