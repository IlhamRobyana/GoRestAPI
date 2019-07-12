package controllers

import (
	models "gorestapi/models"
	"net/http"
	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	result := models.CreateUser(&user)
	return c.String(http.StatusOK, "ok")
}
