package main

import (
	"log"
	"net/http"

	"github.com/Authentication1/routes"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {

	//utils.MigrateDB()

	r := mux.NewRouter()

	routes.SetUserRoutes(r)
	routes.SetLoginRoutes(r)

	srv := http.Server{

		Addr:    ":3000",
		Handler: r,
	}

	log.Println("Running on port 3000")

	log.Println(srv.ListenAndServe())

}
