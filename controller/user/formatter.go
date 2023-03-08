// Package user defines the user controller for the application. The user controller handles incoming
// requests related to user data, and is responsible for converting user model data to JSON format
// for use in the user interface.
package user

import (
	"GoAPIfy/model"
	"time"
)

// UserFormat defines the format in which user data is returned to the user interface.
// It contains the user's ID, name, email, verified_at timestamp, creation timestamp, and update timestamp.
type UserFormat struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// UserWithTokenFormat defines the format in which user data is returned to the user interface
// when a token is included. It contains the user's ID, name, email, token, verified_at timestamp,
// creation timestamp, and update timestamp.
type UserWithTokenFormat struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Token      string     `json:"token"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// UserFormatter is a utility function used to convert a user model to the UserFormat struct.
// It takes a user model as input and returns a UserFormat struct containing
// the user's data in the desired format.
func UserFormatter(user model.User) UserFormat {
	return UserFormat{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		VerifiedAt: user.VerifiedAt,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

// UserWithTokenFormatter is a utility function used to convert a user model to the UserWithTokenFormat struct.
// It takes a user model and a JWT token as input and returns a UserWithTokenFormat struct containing
// the user's data in the desired format, including the provided JWT token.
func UserWithTokenFormatter(user model.User, token string) UserWithTokenFormat {
	return UserWithTokenFormat{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Token:      token,
		VerifiedAt: user.VerifiedAt,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

// UserCollectionFormatter is a utility function used to convert a slice of user models to a
// slice of UserFormat structs.
// It takes a slice of user models as input and returns a slice of UserFormat structs,
// each containing the data for one user in the desired format.
func UserCollectionFormatter(users []model.User) []UserFormat {
	var values []UserFormat
	for _, user := range users {
		values = append(values, UserFormatter(user))
	}
	return values
}
