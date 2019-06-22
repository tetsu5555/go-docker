package main

import (
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend/handlers"
)

func main() {
	// Echo instance
			e := echo.New()

	// Middleware
			e.Use(middleware.Logger())
		e.Use(middleware.Recover())

	// Routes
		e.GET("/", handlers.Hello)

	// Start server
		e.Logger.Fatal(e.Start(":1313"))
}
