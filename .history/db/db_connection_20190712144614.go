package db

import (
        "fmt"
        "os"
        "log"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
)

func CreateCon() *gorm.DB {
        env := godotenv.Load()
        if env != nil{
                fmt.Println(env)
        }

        username := os.Getenv("db_user")
        password := os.Getenv("db_pass")
        dbName := os.Getenv("db_name")
        dbHost := os.Getenv("db_host")
        port := os.Getenv("db_port")
        dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, port, username, dbName, password)

        var err error
        DB, err := gorm.Open("postgres",dbURI)
        if err != nil {
                log.Panic(err)
        }

        //defer DB.Close()

        return DB
}
