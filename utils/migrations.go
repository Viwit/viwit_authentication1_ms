package utils

import (
	"github.com/Authentication1/models"
	//"github.com/Viwit/Authentication1/models"
)

func MigrateDB() {
	db := GetConnection()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Login{})

	//defer db.Close()
}
