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
	result := models.CreateUser(*user)
	return c.JSON(http.StatusOK, result)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	result := models.GetUser(id)
	return c.JSON(http.StatusOK, result)
}

func GetAllUsers(c echo.Context) error{
	result := models.GetAllUsers()
	return c.JSON(http.StatusOK, result)
}

func UpdateUsers(c echo.Context) error{
	var user User

	if err:= c.Bind(&user); err != nil {
		return err
	}
	id := c.Param("id")
	result := models.UpdateUser(id, *user)
	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error{

}

