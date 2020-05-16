package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/joypauls/go-rest-docker/mockdb"
)

const serviceName string = "Users"
const serviceVersion string = "0.0.1"
const defaultPort string = ":10000"

// simulating a db with a slice
var users []mockdb.User = mockdb.FetchAllUsers()

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
