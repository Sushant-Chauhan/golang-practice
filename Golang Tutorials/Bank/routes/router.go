package routes

import (
	"BankingApp/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	return router
}
