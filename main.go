package main

import (
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
	db.Create(&Employee{FirstName: "wasuwat", LastName: "lim", Age: 22})

	router := gin.Default()
	api := router.Group("/api")

	api.GET("employee", GetEmployee)
	api.GET("employee/:id", GetOneEmployee)
	router.Run()
}

// List all employees
func GetEmployee(c *gin.Context) {
	var employee []Employee
	employees := db.Find(&employee)

	c.JSON(http.StatusOK, employees.Value)
}

func GetOneEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp Employee
	employee := db.Where("id = ?", id).First(&emp)
	c.JSON(http.StatusOK, employee)
}
