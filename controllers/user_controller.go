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

func GetUser(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	id := mux.Vars(r)["id"]

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

	users := []models.User{}
	db := utils.GetConnection()
	db.Find(&users)
	j, _ := json.Marshal(users)
	utils.SendResponse(w, http.StatusOK, j)
}

func SetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	db := utils.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	user.Reg_date = time.Now()
	err = db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	j, _ := json.Marshal(user)
	utils.SendResponse(w, http.StatusCreated, j)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userFind := models.User{}
	userData := models.User{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()

	db.Find(&userFind, id)

	if userFind.User_id > 0 {
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			utils.SendErr(w, http.StatusBadRequest)
			return
		}
		db.Model(&userFind).Updates(userData)
		j, _ := json.Marshal(userFind)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	id := mux.Vars(r)["id"]
	db := utils.GetConnection()
	db.Find(&user, id)
	if user.User_id > 0 {
		db.Where("user_id = ?", id).Delete(&user)

		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
}
