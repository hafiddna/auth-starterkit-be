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
	Login(user model.User, rememberToken string) (data map[string]interface{}, err error)
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

func (a *authService) Login(user model.User, rememberToken string) (data map[string]interface{}, err error) {
	if !user.IsActive {
		return nil, fmt.Errorf("user is not active")
	}

	roles := make([]string, len(user.Roles))
	permissions := make([]string, 0)
	teamSubs := make([]helper.JwtAuthClaimTeamSub, 0)
	memberOfs := make([]helper.JwtAuthClaimTeamSub, 0)
	for i, role := range user.Roles {
		if role.Name != "" {
			roles[i] = role.Name
		}
		for _, permission := range role.Permissions {
			if helper.ArrayStringContains(permissions, permission.Name) || permission.Name == "" {
				continue
			}
			permissions = append(permissions, permission.Name)
		}
	}

	for i, team := range user.Teams {
		teamSubs = append(teamSubs, helper.JwtAuthClaimTeamSub{
			Sub:         team.ID,
			Roles:       make([]string, len(team.Roles)),
			Permissions: make([]string, 0),
		})
		for j, role := range team.Roles {
			if role.Name != "" {
				teamSubs[i].Roles[j] = role.Name
			}
			for _, permission := range role.Permissions {
				if helper.ArrayStringContains(teamSubs[i].Permissions, permission.Name) || permission.Name == "" {
					continue
				}
				teamSubs[i].Permissions = append(teamSubs[i].Permissions, permission.Name)
			}
		}
	}

	for i, membersOf := range user.MembersOf {
		memberOfs = append(memberOfs, helper.JwtAuthClaimTeamSub{
			Sub:         membersOf.ID,
			Roles:       make([]string, len(membersOf.Roles)),
			Permissions: make([]string, 0),
		})
		for j, role := range membersOf.Roles {
			if role.Name != "" {
				memberOfs[i].Roles[j] = role.Name
			}
			for _, permission := range role.Permissions {
				if helper.ArrayStringContains(memberOfs[i].Permissions, permission.Name) || permission.Name == "" {
					continue
				}
				memberOfs[i].Permissions = append(memberOfs[i].Permissions, permission.Name)
			}
		}
	}

	authTokenDuration := time.Now().Add(time.Minute * 15)
	authData := helper.JwtAuthClaim{
		TeamSub:     append(teamSubs, memberOfs...),
		Roles:       roles,
		Permissions: permissions,
	}
	authToken := helper.GenerateRS512Token(config.Config.App.JWT.PrivateKey, config.Config.App.Secret.AuthKey, user.ID, authData, authTokenDuration)

	rememberTokenDuration := time.Now().Add(time.Hour * 24)
	rememberData := helper.JwtRememberClaim{
		RememberToken: rememberToken,
	}
	rememberAccessToken := helper.GenerateRS512Token(config.Config.App.JWT.RememberTokenPrivate, config.Config.App.Secret.RememberTokenKey, user.ID, rememberData, rememberTokenDuration)

	return map[string]interface{}{
		"access_token":  authToken,
		"refresh_token": rememberAccessToken,
	}, nil
}

func (a *authService) Profile(id string) (data dto.UserProfileDTO, err error) {
	if !helper.IsUUID(id) {
		return data, fmt.Errorf("invalid id")
	}

	return a.userService.Profile(id)
}
