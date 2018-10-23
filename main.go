package main

import (
	"emp-api/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	// connecting postgres db
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	router := gin.Default()

	api := router.Group("/api")

	api.GET("employee", employee.List)
	api.POST("employee", employee.Add)
	api.PUT("employee/:id", employee.Update)
	api.DELETE("employee/:id", employee.Delete)

	router.Run("localhost:8000")
}
