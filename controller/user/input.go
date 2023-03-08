// Package user defines input structs for user-related requests in the application.
package user

// RegisterInput defines the expected format for request data when registering a new user.
// It contains the user's name, email, password, and confirmation password.
type RegisterInput struct {
	Name      string `json:"name" binding:"required"`      // The user's name (required)
	Email     string `json:"email" binding:"required"`     // The user's email (required)
	Password  string `json:"password" binding:"required"`  // The user's password (required)
	CPassword string `json:"cpassword" binding:"required"` // The user's confirmation password (required)
}

// LoginInput defines the expected format for request data when logging in as a user.
// It contains the user's email and password.
type LoginInput struct {
	Email    string `json:"email"`    // The user's email (required)
	Password string `json:"password"` // The user's password (required)
}
