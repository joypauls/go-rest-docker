package main

import (
	"fmt"
	"log"
	"net/http"
)

const defaultPort string = ":10000"

func healthCheckEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is alive.")
	fmt.Println("Status: Healthy")
}

func handleRequests() {
	http.HandleFunc("/", healthCheckEndpoint)
	log.Fatal(http.ListenAndServe(defaultPort, nil))
}

func main() {
	handleRequests()
}
