package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database! " + err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
