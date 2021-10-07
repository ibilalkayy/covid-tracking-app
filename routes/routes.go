package routes

// Import the libraries
import (
	"net/http"

	"github.com/ibilalkayy/CovidTracker/controllers"
	"github.com/ibilalkayy/CovidTracker/middleware"
)

// Route handles the paths that should be visited
// It uses middleware to check for errors
func Route() {
	http.Handle("/", middleware.ErrorHandling(controllers.Home))
	http.Handle("/signup", middleware.ErrorHandling(controllers.Signup))
	http.Handle("/login", middleware.ErrorHandling(controllers.Login))
	http.Handle("/main", middleware.ErrorHandling(controllers.Main))
	http.Handle("/logout", middleware.ErrorHandling(controllers.Logout))
	// Enabling and to add CSS files
	fileServer := http.FileServer(http.Dir("./views/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
}
