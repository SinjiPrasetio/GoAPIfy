// Package user defines the user controller for the application. The user controller is responsible
// for handling incoming requests related to user data.
package user

import (
	"GoAPIfy/core"
	"GoAPIfy/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler is a struct containing methods for handling user-related requests.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
type UserHandler struct {
	modelService model.Model
}

// NewUserHandler creates a new UserHandler instance and returns a pointer to it.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
func NewUserHandler(modelService model.Model) *UserHandler {
	return &UserHandler{modelService}
}

// CreateUser is a method for handling POST requests related to creating new users.
// It takes a *gin.Context as input and expects the request body to be in JSON format.
// It returns an error response if the input data is invalid or incomplete, or a success
// response if the user is successfully created in the database.
func (h *UserHandler) Create(c *gin.Context) {
	var input RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := core.FormatError(err)
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}
	// Create user in the database and return success response
}

// Login is a method for handling POST requests related to user login.
// It takes a *gin.Context as input and expects the request body to be in JSON format.
// If the input data is invalid or incomplete, it returns an error response.
// If the user credentials are invalid, it returns an unauthorized response.
// If the user credentials are valid, it returns a success response with a JWT token.
func (h *UserHandler) Login(c *gin.Context) {
	// Parse the request body and bind it to a LoginInput struct
	var input LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// If the request body is invalid or incomplete, send an error response
		errorMessage := core.FormatError(err)
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// Validate user credentials and generate a JWT token
	// If the credentials are invalid, send an unauthorized response
	// If the credentials are valid, send a success response with the JWT token
}
