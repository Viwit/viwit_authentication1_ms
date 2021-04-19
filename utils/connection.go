package utils

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func GetConnection() *gorm.DB {
	dsn := "root:jbj4cNRqd7NWnMd@tcp(authentication.cjareiirr0dz.us-east-1.rds.amazonaws.com:3306)/viwit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
