package db

import (
	"backend/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Init() {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "myapp"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{})
	// return db
}

func GetConnection() *gorm.DB {
	return db
}
