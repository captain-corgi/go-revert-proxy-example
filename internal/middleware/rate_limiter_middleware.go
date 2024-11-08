package middleware

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiterMiddleware struct {
	limiter *rate.Limiter
}

func NewRateLimiterMiddleware(rps int, burst time.Duration) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		limiter: rate.NewLimiter(rate.Every(burst/time.Duration(rps)), rps),
	}
}

func (rl *RateLimiterMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
