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

// Create a map to store the status code to message mappings
var statusMessages = map[int]string{
	http.StatusOK:                          "Success to request!",
	http.StatusBadRequest:                  "The request could not be understood or was missing required parameters.",
	http.StatusUnauthorized:                "Authentication failed or user doesn't have permissions for the requested operation.",
	http.StatusForbidden:                   "Access denied. The requested resource requires authentication.",
	http.StatusNotFound:                    "The requested resource was not found.",
	http.StatusMethodNotAllowed:            "The requested method is not supported for the specified resource.",
	http.StatusNotAcceptable:               "The requested resource is not capable of generating a response matching the list of acceptable values.",
	http.StatusConflict:                    "The request conflicts with the current state of the target resource.",
	http.StatusPreconditionFailed:          "The precondition given in the request evaluated to false by the server.",
	http.StatusUnsupportedMediaType:        "The server does not support the media type requested in the request.",
	http.StatusInternalServerError:         "An error occurred on the server.",
	http.StatusServiceUnavailable:          "The server is currently unavailable (overloaded or down).",
	http.StatusContinue:                    "The server has received the request headers and the client should proceed to send the request body.",
	http.StatusSwitchingProtocols:          "The server is switching protocols according to the Upgrade header.",
	http.StatusProcessing:                  "The server is processing the request and will return a response at a later time.",
	http.StatusEarlyHints:                  "The server is sending a hint that a final response is likely to be sent soon.",
	http.StatusCreated:                     "The request has been fulfilled and resulted in a new resource being created.",
	http.StatusAccepted:                    "The request has been accepted for processing, but the processing has not been completed.",
	http.StatusNonAuthoritativeInfo:        "The server is a proxy and received a response from a higher-level server that it believes is authoritative.",
	http.StatusNoContent:                   "The server successfully processed the request and is not returning any content.",
	http.StatusResetContent:                "The server successfully processed the request, but is not returning any content. The client should reset the document view.",
	http.StatusPartialContent:              "The server is delivering only part of the resource due to a range header sent by the client.",
	http.StatusMultiStatus:                 "The server has multiple status codes for the response.",
	http.StatusAlreadyReported:             "The server has already provided the response for this request, and the response has not been modified.",
	http.StatusIMUsed:                      "The server has completed the request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.",
	http.StatusMultipleChoices:             "The client can select a different resource or representation from a list of alternatives.",
	http.StatusMovedPermanently:            "The resource has permanently moved to a different URL.",
	http.StatusFound:                       "The resource has temporarily moved to a different URL.",
	http.StatusSeeOther:                    "The response to the request can be found under a different URL.",
	http.StatusNotModified:                 "The resource has not been modified since the last request.",
	http.StatusUseProxy:                    "The requested resource must be accessed through the proxy given by the Location field.",
	http.StatusTemporaryRedirect:           "The resource resides temporarily under a different URL.",
	http.StatusPermanentRedirect:           "The resource has permanently moved to a different URL.",
	http.StatusPreconditionRequired:        "The server requires the request to be conditional.",
	http.StatusTooManyRequests:             "The user has sent too many requests in a given amount of time.",
	http.StatusRequestHeaderFieldsTooLarge: "The server is unwilling to process the request because its header fields are too large.",
	http.StatusUnavailableForLegalReasons:  "The resource is unavailable due to a legal demand.",
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
	data := FormatResponse(statusMessages[status], status, statusString, response)

	// Send the formatted response object back to the client using the JSON method provided by the Gin framework
	if status == http.StatusOK {
		c.JSON(status, data)
	} else {
		c.AbortWithStatusJSON(status, data)
	}
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
