package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend/db"
	"backend/handlers"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
				e.Use(middleware.Logger())
			e.Use(middleware.Recover())

	// Init DB
	db.Init()

	// Routes
	e.GET("/", handlers.Show)
	e.GET("/get", handlers.Create)

	// Start server
	e.Logger.Fatal(e.Start(":1313"))

}
