package main

import (
	controller "gorestapi/controllers"

	_"github.com/lib/pq"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
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
