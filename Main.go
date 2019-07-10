package main

import (
        "net/http"
        "strconv"
        "log"
	"fmt"
	"database/sql"

        //"github.com/jinzhu/gorm"
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

        db, err := sql.Open("postgres","host=localhost port=5432 user=postgres dbname=restapi_test password=ujangbedil sslmode=disable")
        if err != nil {
                log.Panic(err)
        }

        //defer db.Close()

        user := new(User)
        if err := c.Bind(user); err != nil {
                return err
        }
	fmt.Println(user.First_name)
	age, _ := strconv.ParseInt(user.Age,10,64)
	sqlStatement := "INSERT INTO users (first_name, last_name, age, email) VALUES ($1, $2, $3, $4)"
	res, err := db.Query(sqlStatement, user.First_name, user.Last_name, age, user.Email)
        if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, user)
	}

	//db.Create(&user)
        return c.String(http.StatusOK, "ok")
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
        //e.GET("/users/:id", getUser)
        //e.PUT("/users/:id", updateUser)
        //e.DELETE("/users/:id", deleteUser)

        // Start server
        e.Logger.Fatal(e.Start(":1323"))
}
