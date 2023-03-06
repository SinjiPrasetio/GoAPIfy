package user

import (
	"GoAPI/model"
	"time"
)

type UserFormat struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

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

func UserCollectionFormatter(users []model.User) []UserFormat {
	var values []UserFormat
	for _, user := range users {
		values = append(values, UserFormatter(user))
	}
	return values
}
