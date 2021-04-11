package routes

import (
	"github.com/Authentication1/controllers"

	"github.com/gorilla/mux"
)

func SetUserRoutes(r *mux.Router) {

	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	subRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
}
