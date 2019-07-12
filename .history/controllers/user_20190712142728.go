package controllers

import (
	models "gorestapi/models"
	"net/http"
	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	result := models.CreateUser(user)
	
}
