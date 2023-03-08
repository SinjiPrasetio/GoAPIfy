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
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input RegisterInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := core.FormatError(err)
		core.GiveResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}
	// Create user in the database and return success response
}
