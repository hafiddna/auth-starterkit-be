package service

import (
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type SessionService interface {
	FindOneByAppID(appID string, preloadTokenData bool) (session model.Session, err error)
	Create(session model.Session) error
	Update(session model.Session) error
}

type sessionService struct {
	sessionRepository repository.SessionRepository
}

func NewSessionService(sessionRepository repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: sessionRepository,
	}
}

func (s *sessionService) FindOneByAppID(appID string, preloadTokenData bool) (session model.Session, err error) {
	return s.sessionRepository.FindOneByAppID(appID, preloadTokenData)
}

func (s *sessionService) Create(session model.Session) error {
	return s.sessionRepository.Create(session)
}

func (s *sessionService) Update(session model.Session) error {
	return s.sessionRepository.Update(session)
}
