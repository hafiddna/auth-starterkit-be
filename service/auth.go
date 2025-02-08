package service

import (
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/model"
)

type AuthService interface {
	ValidateUser(dto dto.LoginDTO) (user model.User, err error)
	Login(user model.User) interface{}
	Profile(id string) (data map[string]interface{}, err error)
}

type authService struct {
	userService UserService
}

func NewAuthService(userService UserService) AuthService {
	return &authService{
		userService: userService,
	}
}

func (a *authService) ValidateUser(dto dto.LoginDTO) (user model.User, err error) {
	user, err = a.userService.FindByEmailPhoneOrUsername(dto.Credential)
	if err != nil {
		return user, err
	}

	if helper.ComparePassword(user.Password, dto.Password) {
		return user, nil
	}

	return user, fmt.Errorf("invalid password")
}

func (a *authService) Login(user model.User) interface{} {
	if !user.IsActive {
		return nil
	}

	roles := make([]string, len(user.Roles))
	permissions := make([]string, 0)
	for i, role := range user.Roles {
		roles[i] = role.Name
		for _, permission := range role.Permissions {
			if helper.ArrayStringContains(permissions, permission.Name) {
				continue
			}
			permissions = append(permissions, permission.Name)
		}
	}

	// TODO: Add team_ids to token
	return helper.GenerateToken(user.ID, []string{}, roles, permissions)
}

func (a *authService) Profile(id string) (data map[string]interface{}, err error) {
	if !helper.IsUUID(id) {
		return nil, fmt.Errorf("invalid id")
	}

	return a.userService.Profile(id)
}
