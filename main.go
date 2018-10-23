package main

import (
	"employee-api/handler"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func main() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	fmt.Println(dbURI)

	db, err = gorm.Open("postgres", dbURI)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	router := gin.Default()

	api := router.Group("/api")

	api.GET("employee", employee.List)
	api.POST("employee", employee.Add)
	api.PUT("employee/:id", employee.Update)
	api.DELETE("employee/:id", employee.Delete)

	router.Run("localhost:8000")
}
