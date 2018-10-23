package main

import (
	"emp-api/handler"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func main() {
	// connecting postgres db
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+
			" dbname="+os.Getenv("DBNAME")+" sslmode=disable password="+
			os.Getenv("PASSWORD"))
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
