// Package controller defines the application's controllers, which handle incoming requests
// from the user interface and return responses to the user.
package controller

import (
	"GoAPIfy/controller/user"
	"GoAPIfy/service/appService"
	"GoAPIfy/service/auth"
)

// Handlers defines a struct containing all the application's handlers, each of which
// is responsible for handling a different type of request.
type Handlers struct {
	UserHandler *user.UserHandler // The user handler manages user-related requests
	// Add more handlers as needed
}

// RegisterHandler initializes and returns a struct containing all the application's handlers.
// It takes a model.Model as input, which is used to initialize each of the handlers.
// Returns a pointer to the Handlers struct.
func RegisterHandler(s appService.AppService, authService auth.AuthService) *Handlers {
	return &Handlers{
		UserHandler: user.NewUserHandler(s, authService),
		// Initialize other handlers as needed
	}
}
