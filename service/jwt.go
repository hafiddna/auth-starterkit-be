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
	GenerateToken(userID, teamId, systemRole string, roles, permissions []string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	jwt.RegisteredClaims
	TeamSub     string   `json:"team_sub"`
	Role        string   `json:"role"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type jwtService struct {
	config config.CfgStruct
}

func NewJWTService(config config.CfgStruct) JWTService {
	return &jwtService{
		config: config,
	}
}

func (j *jwtService) GenerateToken(userID, teamId, systemRole string, roles, permissions []string) string {
	uuid := uuid2.New()

	claims := &jwtCustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.App.Name,
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.String(),
		},
		TeamSub:     teamId,
		Role:        systemRole,
		Roles:       roles,
		Permissions: permissions,
	}

	var rsaPrivateKey *rsa.PrivateKey
	var err error

	privateKey := []byte(j.config.App.JWT.PrivateKey)

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
	publicKey := []byte(j.config.App.JWT.PublicKey)

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
