package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

//definir los tipos de datos, user,.....
type Token struct {
	Token_id        int    `json:"token_id,omitempty"`
	Token           string `json:"token,omitempty"`
	Expiration_date string `json:"expiration_id,omitempty"`
	Creation_date   string `json:"creation_date,omitempty"`
}

type Credit_card struct {
	Credit_card_id    int `json:"credit_card_id,omitempty"`
	Credit_card_owner int `json:"credit_card_owner,omitempty"`
	Ccv               int `json:"ccv,omitempty"`
	Expiration_date   int `json:"expiration_id,omitempty"`
}

type User struct {
	User_id       int         `json:"user_id,omitempty"`
	Firstname     string      `json:"firstname,omitempty"`
	Lastname      string      `json:"lastname,omitempty"`
	Email         string      `json:"email,omitempty"`
	Reg_date      string      `json:"reg_date,omitempty"`
	User_password string      `json:"user_password,omitempty"`
	Credit_card   Token       `json:"credit_card,omitempty"`
	Token_id      Credit_card `json:"token_id,omitempty"`
}

type Login struct {
	Login_id   int  `json:"login_id,omitempty"`
	User_id    User `json:"user_id,omitempty"`
	Login_date int  `json:"login_date,omitempty"`
}

//var users []User

func getUserEndpoint(w http.ResponseWriter, req *http.Request) {
	print("por aqui  paso")
	//var usuario User = User{User_id: 1, Firstname: "Jairo", Lastname: "Suarez", Email: "andres4005@gmail.com", Reg_date: "1234", User_password: "admin", Credit_card: 1111, Token_id: 12345}
	json.NewEncoder(w).Encode(req.Body)
}
func setUserEndpoint(w http.ResponseWriter, req *http.Request) {
	print("paso por usuarios")
}
func updateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	print("paso por usuarios")
}
func deleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	print("paso por usuarios")
}

func getloginEndpoint(w http.ResponseWriter, req *http.Request) {
	print("por aqui  paso")
	//var usuario User = User{User_id: 1, Firstname: "Jairo", Lastname: "Suarez", Email: "andres4005@gmail.com", Reg_date: "1234", User_password: "admin", Credit_card: 1111, Token_id: 12345}
	json.NewEncoder(w).Encode(req.Body)
}
func setLoginEndpoint(w http.ResponseWriter, req *http.Request) {
	print("paso por usuarios")
}
func updateLoginEndpoint(w http.ResponseWriter, req *http.Request) {
	print("paso por usuarios")
}
func deleteLoginEndpoint(w http.ResponseWriter, req *http.Request) {
	print("paso por usuarios")
}

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:33060)/viwit")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfuly Connected to MYSQL DB")

	router := mux.NewRouter()

	//users = append(users, User{user_id: "1", firstname: "Jairo", lastname: "Suarez", email: "andres4005@gmail.com", reg_date: "1234", user_password: "admin", credit_card: 1111, token_id: 12345})

	//endopoints
	router.HandleFunc("/user", getUserEndpoint).Methods("POST")
	router.HandleFunc("/user", getUserEndpoint).Methods("GET")
	router.HandleFunc("/user", getUserEndpoint).Methods("PUT")
	router.HandleFunc("/user", getUserEndpoint).Methods("DELETE")

	router.HandleFunc("/login", getUserEndpoint).Methods("POST")
	router.HandleFunc("/login", getUserEndpoint).Methods("GET")
	router.HandleFunc("/login", getUserEndpoint).Methods("PUT")
	router.HandleFunc("/login", getUserEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

}
