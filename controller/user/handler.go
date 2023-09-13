// Package user defines the user controller for the application. The user controller is responsible
// for handling incoming requests related to user data.
package user

import (
	"GoAPIfy/core"
	"GoAPIfy/core/math"
	"GoAPIfy/model"
	"GoAPIfy/service/appService"
	"GoAPIfy/service/auth"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler is a struct containing methods for handling user-related requests.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
type UserHandler struct {
	s           appService.AppService
	authService auth.AuthService
}

// NewUserHandler creates a new UserHandler instance and returns a pointer to it.
// It takes a model.Model as input, which is used to interact with the database and perform
// CRUD operations on user data.
func NewUserHandler(s appService.AppService, authService auth.AuthService) *UserHandler {
	return &UserHandler{s, authService}
}

// CreateUser is a method for handling POST requests related to creating new users.
// It takes a *gin.Context as input and expects the request body to be in JSON format.
// It returns an error response if the input data is invalid or incomplete, or a success
// response if the user is successfully created in the database.
func (h *UserHandler) Register(c *gin.Context) {
	// Parse the request body and bind it to a RegisterInput struct
	var input RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// If the request body is invalid or incomplete, send an error response
		errorMessage := core.FormatError(err)
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
	}

	// Check that the password and confirmation password match
	if input.Password != input.CPassword {
		errorMessage := core.FormatError(errors.New("passwords do not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
	}

	// Hash the password
	hashedPassword, err := math.Hash(input.Password)
	if err != nil {
		errorMessage := core.FormatError(errors.New("failed to hash password"))
		core.SendResponse(c, http.StatusInternalServerError, errorMessage)
	}

	// Create a new User instance with the input data and hashed password
	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// Create the user in the database
	err = h.s.Model.Load(&user).Save()
	if err != nil {
		errorMessage := core.FormatError(errors.New("failed to create user"))
		core.SendResponse(c, http.StatusInternalServerError, errorMessage)
	}

	// Generate a JWT token using the user data
	jwt, err := h.authService.GenerateToken(*user)
	if err != nil {
		errorMessage := core.FormatError(errors.New("failed to generate token"))
		core.SendResponse(c, http.StatusInternalServerError, errorMessage)
	}

	// Format the user data and token into a response object
	response := UserWithTokenFormatter(*user, jwt)

	// Send a success response with the response object
	core.SendResponse(c, http.StatusOK, response)
}

// Login is a method for handling POST requests related to user login.
// It takes a *gin.Context as input and expects the request body to be in JSON format.
// If the input data is invalid or incomplete, its an error response.
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
	}

	email := input.Email       // Get the email from the input data
	password := input.Password // Get the password from the input data

	// Retrieve the user data from the database using the email address as the key
	var userData model.User
	err = h.s.Model.Load(&userData).Where("email = ?", email).Get()
	if err != nil {
		// If there is an error retrieving the user data, send an error response
		// Note: this assumes that the FindByKey meths an error when the key is not found in the database
		errorMessage := core.FormatError(errors.New("email not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
	}

	// Check that the email in the retrieved user data matches the email provided in the login input
	if userData.Email != email {
		errorMessage := core.FormatError(errors.New("email not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
	}

	// Check that the password provided in the login input matches the password in the retrieved user data

	challenge := math.HashChallenge(password, userData.Password)
	if !challenge {
		errorMessage := core.FormatError(errors.New("password not match"))
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// Generate a JWT token using the user data
	jwt, err := h.authService.GenerateToken(userData)
	if err != nil {
		errorMessage := core.FormatError(err)
		core.SendResponse(c, http.StatusUnprocessableEntity, errorMessage)
		return
	}
	// Format the user data and token into a response object
	response := UserWithTokenFormatter(userData, jwt)

	// Send a success response with the response object
	core.SendResponse(c, http.StatusOK, response)
}

// VerifyToken verifies the JWT token from the Authorization header.
// It sends a success response if the token is valid and an error response if the token is invalid or expired.
func (h *UserHandler) VerifyToken(c *gin.Context) {
	// User is set in the context in your middleware
	currentUser, _ := c.Get("currentUser")
	token := c.MustGet("token").(string)

	if currentUser == nil {
		core.SendResponse(c, http.StatusUnauthorized, "Invalid token")
		return
	}
	user := currentUser.(model.User)

	// Send a success response if the token is valid
	userFormat := UserWithTokenFormatter(user, token)
	core.SendResponse(c, http.StatusOK, userFormat)
}

func (h *UserHandler) IsEmailAvailable(c *gin.Context) {
	var input IsEmailAvailableInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := core.FormatError(err)
		core.SendResponse(c, http.StatusInternalServerError, errorMessage)
		return
	}

	var user model.User
	response := false
	err = h.s.Model.Load(&user).Where("email = ?", input.Email).Get()
	if err != nil {
		response = true
	}

	if user.ID == 0 {
		response = true
	}

	core.SendResponse(c, http.StatusOK, response)
}
