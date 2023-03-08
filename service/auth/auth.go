package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"GoAPIfy/model"
)

// AuthService is an interface for handling JWT token generation and validation.
type AuthService interface {
	GenerateToken(user model.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// JWTService is a concrete implementation of the AuthService interface.
type JWTService struct {
	SigningKey []byte
}

// NewJWTService creates a new instance of the JWTService.
func NewJWTService() *JWTService {
	return &JWTService{
		SigningKey: []byte(os.Getenv("APP_KEY")),
	}
}

// GenerateToken generates a JWT token based on user data.
func (s *JWTService) GenerateToken(user model.User) (string, error) {
	// Create a new JWT token with a UUID as the unique identifier (jti)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":    uuid.New().String(),
		"sub":    user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
		"iss":    "GoAPIfy",
		"aud":    "GoAPIfy",
		"nbf":    time.Now().Unix(),
		"iat":    time.Now().Unix(),
		"scopes": []string{"user"},
	})

	// Sign the token with the app key
	signedToken, err := token.SignedString(s.SigningKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates a JWT token and returns the parsed token.
func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	// Parse the JWT token and validate the signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token signing method")
		}
		return s.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, errors.New("Token is invalid")
	}

	return token, nil
}
