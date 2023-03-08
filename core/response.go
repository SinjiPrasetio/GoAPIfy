// Package core provides utility functions and data structures that are commonly used across the GoAPIfy application.
// It contains functions for handling HTTP responses, formatting JSON responses, and extracting error messages from validator.ValidationErrors objects.
// It also includes the Response and Meta data structures, which are used to represent the structure of JSON responses.
package core

import (
	"net/http"

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
// message: string message to be included in the response
// code: HTTP status code to be included in the response
// status: string indicating success or failure status of the response
// data: interface{} representing the data to be included in the response
// returns a formatted Response struct with the provided message, code, status, and data
func FormatResponse(message string, code int, status string, data interface{}) Response {
	// Create a new Meta struct with the given message, code, and status
	jsonMeta := Meta{Message: message, Code: code, Status: status}

	// Create a new Response struct with the Meta struct and the given data
	jsonResponse := Response{Meta: jsonMeta, Data: data}

	return jsonResponse
}

// SendResponse is a utility function that takes in a Gin context object, an HTTP status code, and an interface type representing the response data.
// The function uses the FormatResponse function to create a formatted response object with the provided status code and message.
// The response data is included in the response object as the value of the "data" key.
// Finally, the function sends the response object back to the client using the JSON method provided by the Gin framework.
// c: a Gin context object
// status: HTTP status code to be included in the response
// response: interface{} representing the data to be included in the response
// returns nothing, but sends a formatted response object back to the client using the JSON method provided by the Gin framework
func SendResponse(c *gin.Context, status int, response interface{}) {
	// Determine the status string based on the provided status code
	statusString := http.StatusText(status)

	// Create a formatted response object using the provided message, status code, and response data
	data := FormatResponse(statusString, status, statusString, response)

	// Send the formatted response object back to the client using the JSON method provided by the Gin framework
	c.JSON(status, data)
}

// FormatError takes in an error object and returns a Gin H object with an "errors" key containing a slice of error messages.
// If the error object is a validator.ValidationErrors object, it will extract the error messages from each error in the object.
// If the error object is not a validator.ValidationErrors object, it will simply return a Gin H object with a single error message.
// err: the error object to be processed
// returns a Gin H object with an "errors" key containing a slice of error messages
func FormatError(err error) gin.H {
	// Create a slice to hold error messages
	var errors []string

	// Use type assertion to check if the error is a validator.ValidationErrors object
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// If it is, loop through each error in the object and append its error message to the errors slice
		for _, e := range validationErrors {
			errors = append(errors, e.Error())
		}
	} else {
		// If it's not a validator.ValidationErrors object, simply append the error message to the errors slice
		errors = append(errors, err.Error())
	}

	// Create a Gin H object with an "errors" key containing the errors slice
	return gin.H{"errors": errors}
}
