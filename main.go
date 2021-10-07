package main

// Import the libraries
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ibilalkayy/CovidTracker/routes"
)

// main, calls the route function and sets the port
func main() {
	routes.Route()
	fmt.Println("Starting the server at port: 8080")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
