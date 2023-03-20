package seeder

import (
	"GoAPIfy/factory"
	"GoAPIfy/model"
	"GoAPIfy/service/appService"
)

func RegisterSeeders(s appService.AppService) {

	// Check if the User model has any data in the database
	if count, err := s.Model.Load(&model.User{}).Count(); err != nil {
		panic(err)
	} else if count == 0 {
		// Create a new instance of your UserFactory
		userFactory := &factory.UserFactory{}

		// Create a new instance of your UserSeeder
		userSeeder := NewUserSeeder(s, userFactory)

		// Call the Seed function on your UserSeeder to seed the database with 10 User models
		if err := userSeeder.Seed(10); err != nil {
			panic(err)
		}
	}

	// Repeat the above process for each of your other seeders...
}
