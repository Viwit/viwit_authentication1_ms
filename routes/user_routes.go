package routes

import (
	"github.com/Authentication1/controllers"

	"github.com/gorilla/mux"
)

func SetUserRoutes(r *mux.Router) {

	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	subRouter.HandleFunc("/user/ldap/{idUser}/{password}", controllers.GetUserLdap).Methods("GET")
	subRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	subRouter.HandleFunc("/user", controllers.SetUser).Methods("POST")
	subRouter.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	subRouter.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	subRouter.HandleFunc("/userLogin/", controllers.UserLogin).Methods("GET")
}
