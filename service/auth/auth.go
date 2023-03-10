package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"GoAPIfy/model"
)

// AuthService defines the interface for generating and validating JWT tokens.
type AuthService interface {
	GenerateToken(user model.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// JWTService is the implementation of the AuthService interface.
type JWTService struct {
	SigningKey []byte
}

// NewJWTService creates a new instance of the JWTService.
func NewJWTService() *JWTService {
	return &JWTService{
		SigningKey: []byte(os.Getenv("APP_KEY")),
	}
}

// GenerateToken generates a new JWT token for the provided user.
func (s *JWTService) GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"jti":    uuid.New().String(),
		"sub":    user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iss":    os.Getenv("APP_NAME"),                 // Issuer of the token
		"aud":    os.Getenv("APP_NAME"),                 // Audience for the token
		"nbf":    time.Now().Unix(),                     // Token not valid before this time
		"iat":    time.Now().Unix(),                     // Token issued at this time
		"scopes": []string{"user"},                      // Scopes for the token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.SigningKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateToken validates the provided JWT token and returns the parsed token.
func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	parser := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}
	token, err := parser.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
