package router

import (
	"github.com/gorilla/mux"

	"go-postgres/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users", middleware.GetAllUsers).Methods("GET", "OPTIONS")

	return router
}
