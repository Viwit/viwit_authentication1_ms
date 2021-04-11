package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Authentication1/models"
	"github.com/Authentication1/utils"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	id := mux.Vars(r)["User_id"]

	db := utils.GetConnection()
	//defer db.Close()

	db.Find(&user, id)

	if user.User_id > 0 {
		j, _ := json.Marshal(user)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	println("paso por users :D")

	user := models.User{}
	db := utils.GetConnection()

	db.First(&user)

	j, _ := json.Marshal(user)

	utils.SendResponse(w, http.StatusOK, j)
}
