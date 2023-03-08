// Package user defines the user controller for the application. The user controller is responsible
// for handling incoming requests related to user data.
package user

import (
	"GoAPIfy/core"
	"GoAPIfy/model"
	"GoAPIfy/service/auth"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler is a struct containing methods for handling user-related requests.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
type UserHandler struct {
	modelService model.Model
	authService  auth.Service
}

// NewUserHandler creates a new UserHandler instance and returns a pointer to it.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
func NewUserHandler(modelService model.Model, authService auth.Service) *UserHandler {
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

	// Validate user credentials and generate a JWT token
	email := input.Email
	password := input.Password

	var userModel model.User
	result, err := h.modelService.FindByKey("email", email, &userModel)
	if err != nil {
		// handle error
	}

	fmt.Printf("Type of result before type assertion: %T\n", result)

	userData, ok := result.(*model.User)
	if !ok {
		fmt.Printf("Type assertion failed. Type of result after type assertion: %T\n", result)
		// handle type assertion error
	}

	fmt.Printf("Type of result after type assertion: %T\n", userData)
	if !ok {
		// handle type assertion error
		errorMessage := core.FormatError(errors.New("model not compactible"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	if userData.Email != email {
		errorMessage := core.FormatError(errors.New("email not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		errorMessage := core.FormatError(errors.New("password not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// Create a JWT token
	// You would typically include more information in the token, such as the user ID or other user details
	// For the sake of this example, we'll just create a token with the email as the payload
	jwt, err := h.authService.GenerateToken(*userData)

	response := UserWithTokenFormatter(*userData, jwt)

	// Send a success response with the JWT token
	core.SendResponse(c, http.StatusOK, response)
}
