package core

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response represents the structure of the JSON response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta contains metadata about the response
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// FormatResponse creates a formatted Response struct for the JSON response
func FormatResponse(message string, code int, status string, data interface{}) Response {
	// Create a new Meta struct with the given message, code, and status
	jsonMeta := Meta{Message: message, Code: code, Status: status}

	// Create a new Response struct with the Meta struct and the given data
	jsonResponse := Response{Meta: jsonMeta, Data: data}

	return jsonResponse
}

// FormatValidationErrors extracts error messages from a validator.ValidationErrors object and returns them in a slice
func FormatValidationErrors(err error) []string {
	var errorMessages []string
	var validationErrors validator.ValidationErrors

	// Check if the error is of type validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		// If so, loop through each error in the object and append its error message to the slice
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, e.Error())
		}
	} else {
		// If not, simply append the error message to the slice
		errorMessages = append(errorMessages, err.Error())
	}

	return errorMessages
}

// GiveResponse is a utility function that takes in a Gin context object, an HTTP status code, and an interface type representing the response data.
// The function uses the FormatResponse function to create a formatted response object with the provided status code and message.
// The response data is included in the response object as the value of the "data" key.
// Finally, the function sends the response object back to the client using the JSON method provided by the Gin framework.

func GiveResponse(c *gin.Context, status int, response interface{}) {
	// Create a formatted response object using the provided message, status code, and response data
	data := FormatResponse("Failed to register", status, "error", response)

	// Send the formatted response object back to the client using the JSON method provided by the Gin framework
	c.JSON(status, data)

	return
}
