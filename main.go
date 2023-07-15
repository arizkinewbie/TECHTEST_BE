package main

import (
	"TECHTEST_BE/database"
	"TECHTEST_BE/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the main page of the server. Please use the correct routes."))}).Methods("GET")
	router.HandleFunc("/cakes", handlers.GetCakesHandler).Methods("GET")
	router.HandleFunc("/cakes/{id}", handlers.GetCakeIDHandler).Methods("GET")
	router.HandleFunc("/cakes", handlers.AddCakeHandler).Methods("POST")
	router.HandleFunc("/cakes/{id}", handlers.UpdateCakeHandler).Methods("PUT")
	router.HandleFunc("/cakes/{id}", handlers.DeleteCakeHandler).Methods("DELETE")

	// Start the server
	log.Println("Server started on http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
