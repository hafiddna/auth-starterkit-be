package service

import (
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/model"
	"time"
)

type AuthService interface {
	ValidateUser(dto dto.LoginDTO) (user model.User, err error)
	Login(user model.User) (data map[string]interface{}, err error)
	Profile(id string) (data dto.UserProfileDTO, err error)
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

func (a *authService) Login(user model.User) (data map[string]interface{}, err error) {
	helper.JSONPrettyLog(user)
	if !user.IsActive {
		return nil, fmt.Errorf("user is not active")
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

	authTokenDuration := time.Now().Add(time.Minute * 15)
	// TODO: Add team_ids to token
	authData := helper.JwtAuthClaim{
		TeamSub:     nil,
		Roles:       roles,
		Permissions: permissions,
	}
	authToken := helper.GenerateRS512Token(config.Config.App.JWT.PrivateKey, config.Config.App.Secret.AuthKey, user.ID, authData, authTokenDuration)

	rememberTokenDuration := time.Now().Add(time.Hour * 24)
	rememberData := helper.JwtRememberClaim{
		RememberToken: user.RememberToken.String,
	}
	rememberToken := helper.GenerateRS512Token(config.Config.App.JWT.RememberTokenPrivate, config.Config.App.Secret.RememberTokenKey, user.ID, rememberData, rememberTokenDuration)

	return map[string]interface{}{
		"access_token":  authToken,
		"refresh_token": rememberToken,
	}, nil
}

func (a *authService) Profile(id string) (data dto.UserProfileDTO, err error) {
	if !helper.IsUUID(id) {
		return data, fmt.Errorf("invalid id")
	}

	return a.userService.Profile(id)
}
