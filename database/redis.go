package database

// Import the libraries
import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/go-redis/redis"
)

// Structure to return a token into JSON format
type MyToken struct {
	Token string
}

// Set, create a redis database with a paremeter as key and value
func Set(id, tokenString string) {
	// This variable saves the redis credentials
	client := redis.NewClient(&redis.Options{
		Addr:     "host:port",
		Password: "password",
		DB:       0,
	})

	// It converts the token into JSON
	json, err := json.Marshal(MyToken{Token: tokenString})
	if err != nil {
		log.Fatal(err)
	}

	// It sets the token
	err = client.Set(id, json, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

// Get, fetch the data from redis
func Get(id string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "host:port",
		Password: "password",
		DB:       0,
	})

	// "val" gets the token value from a key id
	val, err := client.Get(id).Result()
	if err != nil {
		return false
	}

	// It removes the unnecessary string using regex
	re, err := regexp.Compile(`.*:"|".*`)
	if err != nil {
		log.Fatal(err)
	}
	_ = re.ReplaceAllString(val, "")
	return true
}

// Del removes the data from redis
func Del(id string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "host:port",
		Password: "password",
		DB:       0,
	})

	// It removes the data by giving a key id
	_, err := client.Del(id).Result()
	if err != nil {
		log.Fatal(err)
	}
}
