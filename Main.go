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

func createUser(c echo.Context) error {

        db, err := gorm.Open("postgres","host=localhost port=5432 user=postgres dbname=restapi_test password=ujangbedil sslmode=disable")
        if err != nil {
                log.Panic(err)
        }

        defer db.Close()

        user := new(User)
        if err := c.Bind(user); err != nil {
                return err
        }
	db.Create(&user)
        return c.String(http.StatusOK, "ok")
}

func getUser(c echo.Context) error {
	db, err := gorm.Open("postgres","host=localhost port=5432 user=postgres dbname=restapi_test password=ujangbedil sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	var user User
	id := c.Param("id")
	db.Where("id=?", id).Find(&user)
	return c.JSON(http.StatusOK, user)
}
func main() {

        //db, err := gorm.Open("postgres","host=localhost port=5432 user=postgres dbname=test password=ujangbedil sslmode=disable")
        //if err != nil {
          //      log.Panic(err)
        //}

        //defer db.Close()
        //db.AutoMigrate(&User{})

        e := echo.New()

        // Middleware
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())

        // Routes
        e.POST("/users", createUser)
        e.GET("/users/:id", getUser)
        //e.PUT("/users/:id", updateUser)
        //e.DELETE("/users/:id", deleteUser)

        // Start server
        e.Logger.Fatal(e.Start(":1323"))
}
