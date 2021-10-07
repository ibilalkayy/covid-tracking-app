package controllers

// Import the libraries
import (
	"net/http"
	"time"

	"github.com/ibilalkayy/CovidTracker/database"
	"github.com/ibilalkayy/CovidTracker/models"
)

// Home returns the home template
// In case another page is executed, it executes the PageError template.
func Home(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return PageError.Execute(w, nil)
	}
	return HomeTmpl.Execute(w, nil)
}

// Signup returns the signup template and takes the users info for registration
// It uses MethodPost to check if it is really posting the data.
func Signup(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return SignupTmpl.Execute(w, nil)
	}

	// models.User struct stores the form values
	user := models.User{
		Name:            r.FormValue("name"),
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirm-password"),
	}
	// HashPassword converts the password into hashes
	hashPassword := HashPassword([]byte(user.Password))
	// If-statement either compare two password or returns an error
	if ComparePasswords(hashPassword, []byte(user.ConfirmPassword)) {
		stuff := SignupData{user.Name, user.Email, hashPassword}
		// InsertData into the database
		database.InsertData(stuff)
		sm := Messages{SignupSuccess: "Your account has been successfully created!"}
		return LoginTmpl.Execute(w, sm)
	} else {
		fm := Messages{SignupFailure: "Both passwords are not matched"}
		return SignupTmpl.Execute(w, fm)
	}
}

// Login returns the login template
// It uses JWT token, and redis to make the login process possible
func Login(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return LoginTmpl.Execute(w, nil)
	}

	user := models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	// FindAccount finds the given data in the database
	match := database.FindAccount(user.Email, user.Password)
	if match {
		// IsAuthorized generates the JWT token & set the expiration time to 5 minutes
		// It stores the token in a cookie
		// It returns token value and boolean to check the authentication
		tokenString, ok := IsAuthorized(w, r, time.Now().Add(5*time.Minute))
		if ok {
			// Set function setup the id1 as a key and tokenString as a value
			database.Set("id1", tokenString)
			http.Redirect(w, r, "/main", http.StatusFound)
			return nil
		} else {
			fm := Messages{LoginFailure: "Enter the correct email or password"}
			return LoginTmpl.Execute(w, fm)
		}
	} else {
		fm := Messages{LoginFailure: "Enter the correct email or password"}
		return LoginTmpl.Execute(w, fm)
	}
}

// Main returns either the main template or the login template
func Main(w http.ResponseWriter, r *http.Request) error {
	// It gets the generated redis-token from the above function
	match := database.Get("id1")
	if match {
		// It refreshes the JWT token after 5 minutes
		RefreshToken(w, r)
		return MainTmpl.Execute(w, nil)
	} else {
		return LoginTmpl.Execute(w, nil)
	}
}

// Logout returns the logout template
func Logout(w http.ResponseWriter, r *http.Request) error {
	// It deletes the redis token
	database.Del("id1")
	// It deletes the JWT token from the cookie
	_, _ = IsAuthorized(w, r, time.Now().Add(0*time.Minute))
	return HomeTmpl.Execute(w, nil)
}
