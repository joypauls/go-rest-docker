package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

const serviceName string = "Users"
const serviceVersion string = "0.0.1"
const defaultPort string = ":10000"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func generateUsers() []User {
	var users = []User{
		User{ID: "123456789", Name: "Beatrice", Email: "beatrice@gmail.com"},
		User{ID: "987654321", Name: "Amy", Email: "amy@amy.net"},
	}
	return users
}

// simulating a db with a slice
var users []User = generateUsers()

type Health struct {
	Status string `json:"status"`
}

// hardcode for now
var healthy = Health{Status: "OK"}

// health check
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint: /health")
	json.NewEncoder(w).Encode(healthy)
}

// get all "records"
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint: /users")
	json.NewEncoder(w).Encode(users)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/health", healthCheck)
	router.HandleFunc("/users", getAllUsers)
	log.Fatal(http.ListenAndServe(defaultPort, router))
}

func main() {
	fmt.Printf("%s API v%s \n", serviceName, serviceVersion)
	handleRequests()
}
