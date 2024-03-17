package main

import (
	"log"
	"net/http"

	"github.com/adityasshinde/Golang-Backend/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/citizens", handlers.GetCitizens).Methods("GET")
	router.HandleFunc("/api/citizens/{id}", handlers.GetCitizen).Methods("GET")
	router.HandleFunc("/api/citizens", handlers.CreateCitizen).Methods("POST")
	router.HandleFunc("/api/citizens/{id}", handlers.UpdateCitizen).Methods("PUT")
	router.HandleFunc("/api/citizens/{id}", handlers.DeleteCitizen).Methods("DELETE")

	// Enable CORS
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Start server with CORS middleware
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", corsHandler(router))
}
