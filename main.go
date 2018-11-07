package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Employee struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100)"`
	LastName  string `gorm:"type:varchar(100)"`
	Age       uint
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USER")+
			" dbname="+os.Getenv("DB_NAME")+" sslmode=disable password="+os.Getenv("DB_PASS"))

	if err != nil {
		log.Println("err", err)
	}

	defer db.Close()
	db.AutoMigrate(&Employee{})

	router := gin.Default()
	api := router.Group("/api")

	api.GET("employee", GetEmployee)
	api.GET("employee/:id", GetOneEmployee)

	router.Run()
}

func GetEmployee(c *gin.Context) {
	var employees []Employee
	err := db.Find(&employees).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusServiceUnavailable, nil)
	} else {
		c.JSON(http.StatusOK, employees)
	}
}

func GetOneEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp Employee
	err := db.Where("id = ?", id).First(&emp).Error

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, nil)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}
