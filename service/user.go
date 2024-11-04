package service

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type UserService interface {
	FindByEmailPhoneOrUsername(credential string) entity.User
	Profile(id string) map[string]interface{}
}

type userService struct {
	userRepository repository.UserRepository
	userProfile    repository.UserProfileRepository
	roleUser       repository.RoleUserRepository
}

func NewUserService(userRepository repository.UserRepository, userProfile repository.UserProfileRepository, roleUser repository.RoleUserRepository) UserService {
	return &userService{
		userRepository: userRepository,
		userProfile:    userProfile,
		roleUser:       roleUser,
	}
}

func (s *userService) FindByEmailPhoneOrUsername(credential string) entity.User {
	return s.userRepository.FindByEmailPhoneOrUsername(credential)
}

func (s *userService) Profile(id string) map[string]interface{} {
	user := s.userRepository.FindOneById(id)
	profile := s.userProfile.FindOneByUserID(id)

	return map[string]interface{}{
		"username":          user.Username,
		"email_verified_at": user.EmailVerifiedAt,
		"phone_verified_at": user.PhoneVerifiedAt,
		"full_name":         profile.(entity.UserProfile).FullName,
		"nick_name":         profile.(entity.UserProfile).NickName,
	}
}
