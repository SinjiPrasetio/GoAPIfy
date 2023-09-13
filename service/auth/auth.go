package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"GoAPIfy/model"
	"GoAPIfy/service/appService"
)

type AuthService interface {
	GenerateToken(user model.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JWTService struct {
	SigningKey []byte
	s          appService.AppService
}

func NewJWTService(s appService.AppService) *JWTService {
	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if len(signingKey) < 32 { // HS256 requires at least 256 bit key
		panic("JWT_SIGNING_KEY must be at least 32 characters long")
	}

	return &JWTService{
		SigningKey: []byte(signingKey),
		s:          s,
	}
}

func (s *JWTService) GenerateToken(user model.User) (string, error) {
	fmt.Println(string(s.SigningKey))
	claims := jwt.MapClaims{
		"jti":   uuid.New().String(),
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"nbf":   time.Now().Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.SigningKey)
}

func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	parser := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}
	fmt.Println(string(s.SigningKey))
	return parser.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.SigningKey, nil
	})
}
