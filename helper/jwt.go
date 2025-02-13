package helper

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	uuid2 "github.com/google/uuid"
	"github.com/hafiddna/auth-starterkit-be/config"
	"log"
	"time"
)

type jwtGeneralClaim struct {
	jwt.RegisteredClaims
	Data EncryptedData `json:"data"`
}

type JwtAuthClaimTeamSub struct {
	Sub         string   `json:"sub"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type JwtAuthClaim struct {
	TeamSub     []JwtAuthClaimTeamSub `json:"team_sub"`
	Roles       []string              `json:"roles"`
	Permissions []string              `json:"permissions"`
}

type JwtRememberClaim struct {
	RememberToken string `json:"remember_token"`
}

func GenerateRS512Token(privateKey, key, userID string, data interface{}, duration time.Time) string {
	uuid := uuid2.New()

	var claims jwt.Claims
	var err error

	marshalledData := JSONMarshal(data)

	encryptedData, err := EncryptAES256CBC([]byte(marshalledData), []byte(key))
	if err != nil {
		log.Fatalf("Error encrypting data: %v", err)
	}

	claims = &jwtGeneralClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Config.App.ServerName,
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(duration),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.String(),
			Audience: jwt.ClaimStrings{
				config.Config.App.Server.URL,
			},
		},
		Data: EncryptedData{
			IV:    encryptedData.IV,
			Value: encryptedData.Value,
			MAC:   encryptedData.MAC,
			Tag:   encryptedData.Tag,
		},
	}

	var rsaPrivateKey *rsa.PrivateKey

	bytePrivateKey := []byte(privateKey)

	rsaPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytePrivateKey)

	if err != nil {
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	t, err := token.SignedString(rsaPrivateKey)

	if err != nil {
		return ""
	}

	return t
}

func ValidateRS512Token(publicKey, token string) (*jwt.Token, error) {
	bytePublicKey := []byte(publicKey)

	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(bytePublicKey)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
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

func GenerateHS256Token(secret, key, userID string, data interface{}, duration time.Time) string {
	uuid := uuid2.New()

	var claims jwt.Claims
	var err error

	marshalledData := JSONMarshal(data)

	encryptedData, err := EncryptAES256CBC([]byte(marshalledData), []byte(key))
	if err != nil {
		log.Fatalf("Error encrypting data: %v", err)
	}

	claims = &jwtGeneralClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Config.App.ServerName,
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(duration),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.String(),
			Audience: jwt.ClaimStrings{
				config.Config.App.Server.URL,
			},
		},
		Data: EncryptedData{
			IV:    encryptedData.IV,
			Value: encryptedData.Value,
			MAC:   encryptedData.MAC,
			Tag:   encryptedData.Tag,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return ""
	}

	return t
}

func ValidateHS256Token(secret, token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}

		return nil, err
	}

	return parsedToken, nil
}
