package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter() func(http.Handler) http.Handler {
	limiter := rate.NewLimiter(10, 20) // 10 req/sec, burst of 20

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if !limiter.Allow() {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
