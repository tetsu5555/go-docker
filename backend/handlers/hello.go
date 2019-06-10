package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	// "backend/config"
)

// Handler
func Hello(c echo.Context) error {
	// db := config.GetDBConnection()
	return c.String(http.StatusOK, "Hello, World! every one!! how are you?")
}