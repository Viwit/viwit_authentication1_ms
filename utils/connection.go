package utils

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func GetConnection() *gorm.DB {
	dsn := "root:secret@tcp(127.0.0.1:33060)/viwit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
