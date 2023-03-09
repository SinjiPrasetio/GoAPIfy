// Package config provides configuration options for the application.
package config

// AllowOriginConfig returns a slice of strings representing the URLs allowed to make
// cross-origin requests to the application. You can modify the URLs returned by this function
// to allow requests from other domains, as needed for your specific use case.
// By default, this implementation allows requests only from http://localhost:3000.
func AllowOriginConfig() []string {
	return []string{
		"http://localhost:3000",
	}
}
