package main

import (
	"BankingApp/models"
	"BankingApp/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	models.InitializeDB()

	// Initialize routes
	router := routes.InitRoutes()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
