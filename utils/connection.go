package utils

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func GetConnection() *gorm.DB {
	dsn := "ikHV5b7fXF:sub54mrAjY@tcp(remotemysql.com:3306)/ikHV5b7fXF?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

func CloseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.Close()
}
