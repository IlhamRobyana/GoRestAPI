package main

import (
        "net/http"
        "log"
		"fmt"
		"os"

		controller "gorestapi/controllers"

        "github.com/joho/godotenv"
		"github.com/jinzhu/gorm"
		_"github.com/lib/pq"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
)

func main() {
	env := godotenv.Load()
	if env != nil{
		fmt.Println(env)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	port := os.Getenv("db_port")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, port, username, dbName, password)

	var err error
	DB, err = gorm.Open("postgres",dbUri)
	if err != nil {
			log.Panic(err)
	}

	defer DB.Close()
	DB.AutoMigrate(&User{})

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", controller.CreateUser)
	e.GET("/users/:id", controller.GetUser)
	e.GET("/users", controller.GetAllUsers)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
