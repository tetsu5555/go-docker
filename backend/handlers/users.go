package handlers

import (
	"backend/db"
	"backend/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Handler
func Show(c echo.Context) error {
	user := models.User{}

	// TODO: FIX
	user.Name = "hogehoge" + strconv.Itoa(rand.Intn(10000))
	user.Password = "taro"
	user.Hobby = "fishing"

	db := db.GetConnection()
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func Create(c echo.Context) error {
	user := models.User{}
	// user.id = 1
	// user.Name = "hogehoge2"
	// user.Password = "taro"
	// user.Hobby = "fishing"
	db := db.GetConnection()
	err := db.Where("id = ?", 10).Find(&user).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
