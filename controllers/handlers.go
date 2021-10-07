package controllers

// Import the libaries
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ibilalkayy/CovidTracker/database"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// makeTemplate takes the common files used in every template
func makeTemplate(path string) *template.Template {
	files := []string{path, "views/templates/base.html", "views/templates/footer.html"}
	return template.Must(template.ParseFiles(files...))
}

// Multiple templates
var (
	HomeTmpl   = makeTemplate("views/templates/home.html")
	SignupTmpl = makeTemplate("views/templates/signup.html")
	LoginTmpl  = makeTemplate("views/templates/login.html")
	MainTmpl   = makeTemplate("views/templates/main.html")
	PageError  = makeTemplate("views/templates/pageerror.html")
)

// Structure to take signup info
type SignupData struct {
	Name, Email, Password string
}

// Struture to show different messages
type Messages struct {
	SignupSuccess, SignupFailure, LoginFailure string
}

// Structure for the claims
type SomeClaims struct {
	Email, Password string
	jwt.StandardClaims
}

var Claims = &SomeClaims{}

// HashPassword convert password into hash and returns string
func HashPassword(pass []byte) string {
	// It converts the password into hash with the given cost 4
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// ComparePasswords compare two passwords and return boolean
func ComparePasswords(hashPass string, plainPass []byte) bool {
	hashByte := []byte(hashPass)
	// It compares hashbyte with the plain text
	if err := bcrypt.CompareHashAndPassword(hashByte, plainPass); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// LoadEnvVariable stores the secret key in environment variable and returns []byte
func LoadEnvVariable() []byte {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	// Access the secret key
	value, ok := viper.Get("ACCESS_SECRET").(string)
	if !ok {
		log.Fatal(err)
	}
	byteEnv := []byte(value)
	return byteEnv
}

// GenerateJWT generates token and returns string and error
func GenerateJWT(w http.ResponseWriter, r *http.Request) (string, error) {
	// Set the 5 minutes expiration time limit
	expirationTime := time.Now().Add(5 * time.Minute)
	claim := &SomeClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Use secret key in the claims
	byteEnv := LoadEnvVariable()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(byteEnv)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, err
}

// Authorize the JWT token and returns string and boolean
func IsAuthorized(w http.ResponseWriter, r *http.Request, expirationTime time.Time) (string, bool) {
	tokenString, err := GenerateJWT(w, r)
	if err != nil {
		log.Fatal(err)
	}

	// Set the cookie in the browser
	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})

	// Use the secret key in parse claims
	byteEnv := LoadEnvVariable()
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (interface{}, error) {
		return byteEnv, err
	})

	if token.Valid {
		return tokenString, true
	} else {
		return "", false
	}
}

// RefreshToken renews the token after every 5 minutes
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Generate a new token 30 seconds before the expiration time
	if time.Until(time.Now().Add(time.Duration(Claims.ExpiresAt))) > 30*time.Second {
		fmt.Println("Error")
		return
	}

	tokenString, ok := IsAuthorized(w, r, time.Now().Add(5*time.Minute))
	if ok {
		database.Set("id1", tokenString)
	}
}
