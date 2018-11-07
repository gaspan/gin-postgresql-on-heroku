package main

import (
	"employee-api/handler"
	"employee-api/model"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USER")+
			" dbname="+os.Getenv("DB_NAME")+" sslmode=disable password="+os.Getenv("DB_PASS"))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&model.Employee{})

	router := gin.Default()
	api := router.Group("/api")

	api.GET("employee", ListEmployees)
	api.POST("employee", employee.Add)
	api.PUT("employee/:id", employee.Update)
	api.DELETE("employee/:id", employee.Delete)

	router.Run(":5000")
}

func ListEmployees(c *gin.Context) {
	var employee []model.Employee
	if err := db.Find(&employee).Error; err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}
