package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		id := r.Context().Value(RequestIDKey)

		log.Printf(
			"[%v] %s %s %s",
			id,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
