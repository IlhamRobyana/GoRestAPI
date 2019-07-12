package models

import (
	db "gorestapi/db"
	"net/http"
	
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq" //needed
	"github.com/labstack/echo"
)

type User struct {
	Id		string	`json:"id"`
	Age		string	`json:"age"`
	First_name	string	`json:"first_name"`
	Last_name	string	`json:"last_name"`
	Email		string	`json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

var DB *gorm.DB

func CreateUser(user User) error {
	DB := db.CreateCon()
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	DB.Create(&user)
	return c.String(http.StatusOK, "ok")
}

func GetUser(c echo.Context) error {
	DB := db.CreateCon()
	var user User
	id := c.Param("id")
	DB.Where("id=?", id).Find(&user)
	return c.JSON(http.StatusOK, user)
}

func GetAllUsers(c echo.Context) error {
	DB := db.CreateCon()
	var users []User
	DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	DB := db.CreateCon()
	var user User

	if err:= c.Bind(&user); err != nil {
		return err
	}
	id := c.Param("id")
	attrMap := map[string]interface{}{"first_name": user.First_name,"last_name": user.Last_name, "age": user.Age, "email": user.Email}
	DB.Model(&user).Where("id=?", id).Updates(attrMap)
	return c.NoContent(http.StatusOK)
}

func DeleteUser(c echo.Context) error {
	DB := db.CreateCon()
	var user User

	id := c.Param("id")
	DB.Where("id=?",id).Find(&user).Delete(&user)
	return c.JSON(http.StatusOK, user)
}