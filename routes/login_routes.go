package routes

import (
	"github.com/Authentication1/controllers"

	"github.com/gorilla/mux"
)

func SetLoginRoutes(r *mux.Router) {

	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/login/{id}", controllers.GetLogin).Methods("GET")
	subRouter.HandleFunc("/logins", controllers.Getlogins).Methods("GET")
	subRouter.HandleFunc("/login", controllers.SetLogin).Methods("POST")
}
