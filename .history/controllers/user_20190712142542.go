package controllers

import (
	models "gorestapi/models"
	"net/http"
	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	result := models.CreateUser()
}
