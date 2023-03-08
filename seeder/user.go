package seeder

import (
	"GoAPIfy/factory"
	"GoAPIfy/model"
)

type UserSeeder struct {
	ModelService model.Model
	Factory      *factory.UserFactory
}

func NewUserSeeder(modelService model.Model, factory *factory.UserFactory) *UserSeeder {
	return &UserSeeder{
		ModelService: modelService,
		Factory:      factory,
	}
}

func (s *UserSeeder) Seed(count int) error {
	for i := 0; i < count; i++ {
		user, err := s.Factory.Generate("password")
		if err != nil {
			panic(err)
		}

		if _, err := s.ModelService.Create(&user); err != nil {
			return err
		}
	}

	return nil
}
