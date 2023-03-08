// Package user defines the user controller for the application. The user controller is responsible
// for handling incoming requests related to user data.
package user

import (
	"GoAPIfy/core"
	"GoAPIfy/model"
	"GoAPIfy/service/auth"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler is a struct containing methods for handling user-related requests.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
type UserHandler struct {
	modelService model.Model
	authService  auth.AuthService
}

// NewUserHandler creates a new UserHandler instance and returns a pointer to it.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
func NewUserHandler(modelService model.Model, authService auth.AuthService) *UserHandler {
	return &UserHandler{modelService, authService}
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

	email := input.Email       // Get the email from the input data
	password := input.Password // Get the password from the input data

	// Retrieve the user data from the database using the email address as the key
	var userModel model.User
	result, err := h.modelService.FindByKey("email", email, &userModel)
	if err != nil {
		// If there is an error retrieving the user data, send an error response
		// Note: this assumes that the FindByKey method returns an error when the key is not found in the database
		errorMessage := core.FormatError(err)
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// Type assert the result to a pointer to a User model
	userData, ok := result.(*model.User)
	if !ok {
		// If the type assertion fails, send an error response
		errorMessage := core.FormatError(errors.New("model not compactible"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// Check that the email in the retrieved user data matches the email provided in the login input
	if userData.Email != email {
		errorMessage := core.FormatError(errors.New("email not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// Check that the password provided in the login input matches the password in the retrieved user data
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		errorMessage := core.FormatError(errors.New("password not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// Generate a JWT token using the user data
	jwt, err := h.authService.GenerateToken(*userData)

	// Format the user data and token into a response object
	response := UserWithTokenFormatter(*userData, jwt)

	// Send a success response with the response object
	core.SendResponse(c, http.StatusOK, response)
}
