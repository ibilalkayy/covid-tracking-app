package middleware

// Import the libraries
import (
	"log"
	"net/http"
)

// Type to be used as a parameter in a function
type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ErrorHandling handles the error and returns http.Handler
func ErrorHandling(h MyHandlerFunc) http.Handler {
	// The reason for returning a HandlerFunc is to use it as a middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Fatal(err)
		}
	})
}
