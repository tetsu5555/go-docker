package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"backend/handlers"

	"time"
)

func gormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME   := "myapp"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type User struct {
	ID                                     uint64     `json:"id"`
	Identifier                             string     `json:"identifier"              validate:"required,max=15"`
	Name                                   string     `json:"name"                    validate:"required,max=50"`
	Profile                                string     `json:"profile"                 validate:"max=160"`
	Avatar                                 string     `json:"avatar"                  validate:"max=255"`
	CreatedAt                              time.Time  `json:"created_at"`
	UpdatedAt                              time.Time  `json:"updated_at"`
	DeletedAt                              *time.Time `json:"deleted_at"`
}


func main() {
	// Echo instance
	e := echo.New()

	db := gormConnect()
	defer db.Close()

	// 構造体のインスタンス化
	user := User{Name: "name"}
	db.Create(&user)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1313"))
}
