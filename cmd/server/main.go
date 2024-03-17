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

	// Start server
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", router)
}
