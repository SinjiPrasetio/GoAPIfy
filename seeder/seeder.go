package seeder

import (
	"GoAPIfy/factory"
	"GoAPIfy/model"
)

func RegisterSeeders(modelService model.Model) {

	// Check if the User model has any data in the database
	if count, err := modelService.Count(&model.User{}); err != nil {
		panic(err)
	} else if count == 0 {
		// Create a new instance of your UserFactory
		userFactory := &factory.UserFactory{}

		// Create a new instance of your UserSeeder
		userSeeder := NewUserSeeder(modelService, userFactory)

		// Call the Seed function on your UserSeeder to seed the database with 10 User models
		if err := userSeeder.Seed(10); err != nil {
			panic(err)
		}
	}

	// Repeat the above process for each of your other seeders...
}
