package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Authentication1/models"
	"github.com/Authentication1/utils"

	"github.com/gorilla/mux"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {

	login := models.Login{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()

	db.Find(&login, id)

	if login.Login_id > 0 {
		j, _ := json.Marshal(login)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
}

func SetLogin(w http.ResponseWriter, r *http.Request) {
	login := models.Login{}
	db := utils.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	login.Login_date = time.Now()
	err = db.Create(&login).Error
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	j, _ := json.Marshal(login)
	utils.SendResponse(w, http.StatusCreated, j)
}

func Getlogins(w http.ResponseWriter, r *http.Request) {

	logins := []models.Login{}
	db := utils.GetConnection()
	db.Find(&logins)
	j, _ := json.Marshal(logins)
	utils.SendResponse(w, http.StatusOK, j)
}
