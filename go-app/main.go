package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello from Go API Server!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	log.Println("Server Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
