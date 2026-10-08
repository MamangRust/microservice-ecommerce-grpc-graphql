package middlewares

import (
	"net/http"

	"golang.org/x/time/rate"
)

const (
	// defaultRateLimitRPS is the steady-state request budget per second.
	defaultRateLimitRPS = 100
	// defaultRateLimitBurst is how many requests may arrive back-to-back
	// before the limiter starts rejecting.
	defaultRateLimitBurst = 200
)

// RateLimiter wraps a token-bucket limiter so it can guard any net/http
// handler chain (GraphQL included). It was previously an Echo-only middleware
// and is now usable from the gqlgen server.
type RateLimiter struct {
	limiter *rate.Limiter
}

// NewRateLimiter builds a limiter allowing rps requests per second with the
// given burst. A non-positive rps or burst falls back to the defaults below.
func NewRateLimiter(rps int, burst int) *RateLimiter {
	if rps <= 0 {
		rps = defaultRateLimitRPS
	}
	if burst <= 0 {
		burst = defaultRateLimitBurst
	}

	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Limit(rps), burst),
	}
}

// Middleware rejects the request with HTTP 429 once the bucket is empty.
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.limiter.Allow() {
			writeJSONError(w, "too many requests, please try again later", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
