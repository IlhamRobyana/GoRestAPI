package main

import (
        "net/http"
        _ "strconv"
        "log"
	_ "fmt"
	_ "database/sql"

        "github.com/jinzhu/gorm"
        _ "github.com/lib/pq"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
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

func createUser(c echo.Context) error {
        user := new(User)
        if err := c.Bind(user); err != nil {
                return err
        }
	DB.Create(&user)
        return c.String(http.StatusOK, "ok")
}

func getUser(c echo.Context) error {
	var user User
	id := c.Param("id")
	DB.Where("id=?", id).Find(&user)
	return c.JSON(http.StatusOK, user)
}

func getAllUsers(c echo.Context) error {
	var users []User
	DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func updateUser(c echo.Context) error {
	var user User

	if err:= c.Bind(&user); err != nil {
		return err
	}
	id := c.Param("id")
	attrMap := map[string]interface{}{"first_name": user.First_name,"last_name": user.Last_name, "age": user.Age, "email": user.Email}
	DB.Model(&user).Where("id=?", id).Updates(attrMap)
	return c.NoContent(http.StatusOK)
}

func deleteUser(c echo.Context) error {

        var user User

	id := c.Param("id")
	DB.Where("id=?",id).Find(&user).Delete(&user)
	return c.JSON(http.StatusOK, user)
}

func main() {

	var err error
        DB, err = gorm.Open("postgres","host=localhost port=5432 user=postgres dbname=restapi_test password=ujangbedil sslmode=disable")
        if err != nil {
                log.Panic(err)
        }

	defer DB.Close()
        //db.AutoMigrate(&User{})

        e := echo.New()

        // Middleware
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())

        // Routes
        e.POST("/users", createUser)
        e.GET("/users/:id", getUser)
	e.GET("/users", getAllUsers)
        e.PUT("/users/:id", updateUser)
        e.DELETE("/users/:id", deleteUser)

        // Start server
        e.Logger.Fatal(e.Start(":1323"))
}
