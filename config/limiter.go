// Package config provides configuration options for the application.
package config

import (
	"time"

	"github.com/ulule/limiter/v3"
)

// LimiterConfig returns the rate limit configuration for the application.
// The returned value is a struct of type limiter.Rate that includes a time period and limit.
// The time period is set to one minute, and the limit is set to 100 requests per minute.
func LimiterConfig() limiter.Rate {
	return limiter.Rate{
		Period: time.Minute,
		Limit:  100,
	}
}
