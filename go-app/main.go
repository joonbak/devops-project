package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "myapp_http_requests_total",
	Help: "The total number of HTTP requests",
})

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestCounter.Inc()
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello from Go API Server!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	prometheus.MustRegister(requestCounter)
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", handler)

	log.Println("Server Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
