package service

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	uuid2 "github.com/google/uuid"
	"github.com/hafiddna/auth-starterkit-be/config"
	"time"
)

type JWTService interface {
	GenerateToken(userID, systemRole string, teamIds, roles, permissions []string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	jwt.RegisteredClaims
	Role        string   `json:"role"`
	TeamSub     []string `json:"team_sub"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type jwtService struct {
}

func NewJWTService() JWTService {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(userID, systemRole string, teamIds, roles, permissions []string) string {
	uuid := uuid2.New()

	claims := &jwtCustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Config.App.Name,
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.String(),
		},
		TeamSub:     teamIds,
		Role:        systemRole,
		Roles:       roles,
		Permissions: permissions,
	}

	var rsaPrivateKey *rsa.PrivateKey
	var err error

	privateKey := []byte(config.Config.App.JWT.PrivateKey)

	rsaPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS512, claims)
	t, err := token.SignedString(rsaPrivateKey)

	if err != nil {
		return ""
	}

	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	publicKey := []byte(config.Config.App.JWT.PublicKey)

	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return rsaPublicKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}

		return nil, err
	}

	return parsedToken, nil
}
