package main

import (
        "net/http"
        "log"
	"fmt"
	"os"

        "github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq"
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
        e.POST("/users", createUser)
        e.GET("/users/:id", getUser)
	e.GET("/users", getAllUsers)
        e.PUT("/users/:id", updateUser)
        e.DELETE("/users/:id", deleteUser)

        // Start server
        e.Logger.Fatal(e.Start(":1323"))
}
