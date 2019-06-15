package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! every one!! how are you?")
}
