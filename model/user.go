package model

import (
	"time"

	"gorm.io/gorm"
)

// User is the model representing a user
type User struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	AvatarPath string
	VerifiedAt *time.Time
}
