package api

import (
	"banking-app/auth"
	"banking-app/handlers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Protected routes
	apiRoutes := router.PathPrefix("/api").Subrouter()
	apiRoutes.Use(auth.Middleware)
	apiRoutes.HandleFunc("/accounts", handlers.CreateAccountHandler).Methods("POST")
	apiRoutes.HandleFunc("/accounts/{id}", handlers.GetAccountHandler).Methods("GET")
	apiRoutes.HandleFunc("/transfers", handlers.CreateTransferHandler).Methods("POST")
	apiRoutes.HandleFunc("/transfers", handlers.ListTransfersHandler).Methods("GET")

	return router
}
