package service

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type UserService interface {
	FindByEmailPhoneOrUsername(credential string) (user entity.User, err error)
	Profile(id string) (data map[string]interface{}, err error)
}

type userService struct {
	userRepository repository.UserRepository
	userProfile    repository.UserProfileRepository
	userSetting    repository.UserSettingRepository
	roleUser       repository.RoleUserRepository
}

func NewUserService(userRepository repository.UserRepository, userProfile repository.UserProfileRepository, userSetting repository.UserSettingRepository, roleUser repository.RoleUserRepository) UserService {
	return &userService{
		userRepository: userRepository,
		userProfile:    userProfile,
		userSetting:    userSetting,
		roleUser:       roleUser,
	}
}

func (s *userService) FindByEmailPhoneOrUsername(credential string) (user entity.User, err error) {
	return s.userRepository.FindByEmailPhoneOrUsername(credential)
}

func (s *userService) Profile(id string) (data map[string]interface{}, err error) {
	user, err := s.userRepository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	profile, err := s.userProfile.FindOneByUserID(id)
	if err != nil {
		return nil, err
	}

	data = map[string]interface{}{
		"username":          user.Username,
		"email_verified_at": user.EmailVerifiedAt,
		"phone_verified_at": user.PhoneVerifiedAt,
		"full_name":         profile.FullName,
		"nick_name":         profile.NickName,
	}

	return data, nil
}
