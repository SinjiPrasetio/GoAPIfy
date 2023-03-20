// Package seeder provides functionality to seed data in a database using factory and appService.
package seeder

import (
	"GoAPIfy/factory"
	"GoAPIfy/service/appService"
)

// UserSeeder struct holds references to appService and UserFactory.
type UserSeeder struct {
	AppService appService.AppService
	Factory    *factory.UserFactory
}

// NewUserSeeder creates a new instance of UserSeeder.
func NewUserSeeder(s appService.AppService, f *factory.UserFactory) *UserSeeder {
	return &UserSeeder{
		AppService: s,
		Factory:    f,
	}
}

// Seed creates a specified number of user instances using UserFactory's Generate method and inserts them into the database
// using AppService's Model.Create method.
func (s *UserSeeder) Seed(count int) error {
	for i := 0; i < count; i++ {
		user, err := s.Factory.Generate("password")
		if err != nil {
			panic(err)
		}

		if err := s.AppService.Model.Load(&user).Save(); err != nil {
			return err
		}
	}
	return nil
}
