package main

import (
        "net/http"
        //"strconv"
        "log"


        "github.com/jinzhu/gorm"
        _ "github.com/lib/pq"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
)

type User struct {
        gorm.Model
        id string `json:"id"`
        age int `json:"age"`
        first_name string `json:"first_name"`
        last_name string `json:"last_name"`
        email string `json:"email"`
}

func createUser(c echo.Context) error {

        db, err := gorm.Open("postgres","host=localhost port=5432 user=postgres dbname=test password=ujangbedil sslmode=disable")
        if err != nil {
                log.Panic(err)
        }

        defer db.Close()

        user := new(User)
        if err := c.Bind(user); err != nil {
                return err
        }

        db.Create(&user)
        return c.JSON(http.StatusCreated, user)
}
func main() {

        db, err := gorm.Open("postgres","host=localhost port=5432 user=postgres dbname=test password=ujangbedil sslmode=disable")
        if err != nil {
                log.Panic(err)
        }

        defer db.Close()
        db.AutoMigrate(&User{})

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
