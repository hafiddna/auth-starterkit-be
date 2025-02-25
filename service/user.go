package service

import (
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/hafiddna/auth-starterkit-be/repository"
)

type UserService interface {
	FindByEmailPhoneOrUsername(credential string) (user model.User, err error)
	FindByIDWithTokenData(id string) (user model.User, err error)
	Profile(id string) (data dto.UserProfileDTO, err error)
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

func (s *userService) FindByEmailPhoneOrUsername(credential string) (user model.User, err error) {
	return s.userRepository.FindByEmailPhoneOrUsername(credential)
}

func (s *userService) FindByIDWithTokenData(id string) (user model.User, err error) {
	return s.userRepository.FindByIDWithTokenData(id)
}

func (s *userService) Profile(id string) (data dto.UserProfileDTO, err error) {
	user, err := s.userRepository.FindOneById(id)
	if err != nil {
		return data, err
	}

	profile, err := s.userProfile.FindOneByUserID(id)
	if err != nil {
		return data, err
	}

	data = dto.UserProfileDTO{
		Username:        user.Username.String,
		Email:           user.Email.String,
		EmailVerifiedAt: user.EmailVerifiedAt.Int64,
		Phone:           user.Phone.String,
		PhoneVerifiedAt: user.PhoneVerifiedAt.Int64,
		FullName:        *profile.FullName,
		NickName:        *profile.NickName,
		Role:            user.Roles[0].Name,
		// TODO: Fill this data
		//Avatar:          "",
		//Teams:           nil,
	}

	return data, nil
}
