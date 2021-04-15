package utils

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func GetConnection() *gorm.DB {
	dsn := "uynepu7prokjcn8k:evR43cetezZNG3oaYhW5@tcp(b7hvkzeot7rrwgjxe7qq-mysql.services.clever-cloud.com:3306)/b7hvkzeot7rrwgjxe7qq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
