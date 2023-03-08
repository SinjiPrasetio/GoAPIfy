package controller

import (
	"GoAPIfy/controller/user"
	"GoAPIfy/model"
)

type Handlers struct {
	UserHandler *user.UserHandler
	// Add more handlers as needed
}

func RegisterHandler(modelService model.Model) *Handlers {
	return &Handlers{
		UserHandler: user.NewUserHandler(modelService),
		// Initialize other handlers as needed
	}
}
