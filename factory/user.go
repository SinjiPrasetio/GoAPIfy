package factory

import (
	"Laravel/model"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
)

// UserFactory is a factory that generates User models with randomized data.
type UserFactory struct{}

// Generate generates a new User model with randomized data.
func (f *UserFactory) Generate(password string) (*model.User, error) {
	// Generate a random name and email address
	name := faker.Name()
	email := faker.Email()

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Return a new User instance with the randomized data
	return &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}
