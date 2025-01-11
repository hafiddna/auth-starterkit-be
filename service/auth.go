package service

import (
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/helper"
)

type AuthService interface {
	ValidateUser(dto dto.LoginDto) (user entity.User, err error)
	Login(user entity.User) interface{}
	Profile(id string) (data map[string]interface{}, err error)
}

type authService struct {
	userService UserService
	jwtService  JWTService
}

func NewAuthService(userService UserService, jwtService JWTService) AuthService {
	return &authService{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (a *authService) ValidateUser(dto dto.LoginDto) (user entity.User, err error) {
	user, err = a.userService.FindByEmailPhoneOrUsername(dto.Credential)
	if err != nil {
		return user, err
	}

	if helper.ComparePassword(user.Password, dto.Password) {
		return user, nil
	}

	return user, fmt.Errorf("invalid password")
}

func (a *authService) Login(user entity.User) interface{} {
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

	// TODO:
	//return a.jwtService.GenerateToken(user.ID, user.Role, []string{}, roles, permissions)
	return a.jwtService.GenerateToken(user.ID, "", []string{}, roles, permissions)
}

func (a *authService) Profile(id string) (data map[string]interface{}, err error) {
	if !helper.IsUUID(id) {
		return nil, fmt.Errorf("invalid id")
	}

	return a.userService.Profile(id)
}
