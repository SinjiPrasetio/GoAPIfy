package auth

import (
	"GoAPIfy/model"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(user model.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(user model.User) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = user.ID
	claim["email"] = user.Email
	claim["name"] = user.Name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	verify, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token.")
		}

		return []byte(os.Getenv("APP_KEY")), nil
	})

	if err != nil {
		return verify, err
	}

	return verify, nil
}
