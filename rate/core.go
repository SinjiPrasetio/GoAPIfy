// Package rate provides a rate limiter implementation using the Ulule limiter library.
package rate

import (
	"Laravel/config"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// The NewLimiter function creates a new instance of a rate limiter with configuration parameters from the LimiterConfig
// function in the config package. The rate limiter is implemented using the memory store driver from the Ulule limiter
// library.
//
// Example usage:
// limiter := limiter.NewLimiter()
func NewLimiter() *limiter.Limiter {
	return limiter.New(memory.NewStore(), config.LimiterConfig())
}
