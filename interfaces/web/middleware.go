package web

import (
	"log"
	"net/http"
)

// Middleware is a function type for defining HTTP middleware.
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware is an example middleware that logs incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform logging or request tracking here
		// For example, you can log the request method and URL
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware is an example middleware for authentication.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform authentication checks here
		// For example, you can check for a valid token or session

		// If authentication fails, you can return an error response
		// Otherwise, call the next handler in the chain
		if !IsAuthenticated(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// IsAuthenticated is a placeholder function to check authentication.
func IsAuthenticated(r *http.Request) bool {
	// Implement your authentication logic here
	// For example, check for a valid token or session
	return true // Replace with actual authentication logic
}
