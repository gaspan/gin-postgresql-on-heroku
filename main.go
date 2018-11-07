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

// Employee struct provides basic employee information
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
	api.POST("employee", CreateEmployee)
	api.PUT("employee/:id", UpdateEmployee)
	api.DELETE("employee/:id", DeleteEmployee)

	router.Run()
}

func GetEmployee(c *gin.Context) {
	var emps []Employee
	err := db.Find(&emps).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusServiceUnavailable, nil)
	} else {
		c.JSON(http.StatusOK, emps)
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

func CreateEmployee(c *gin.Context) {
	var emp Employee
	err := c.BindJSON(&emp)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		db.Create(&emp)
		c.JSON(http.StatusOK, emp)
	}
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp Employee

	err := db.Where("id = ?", id).First(&emp).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.BindJSON(&emp)
		db.Save(&emp)
		c.JSON(http.StatusOK, emp)
	}
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp Employee
	err := db.Where("id = ?", id).First(&emp).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		db.Delete(&emp)
		c.JSON(http.StatusOK, "success")
	}
}
