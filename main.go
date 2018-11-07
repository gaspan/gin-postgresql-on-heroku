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

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Employee `json:"data"`
}

type ArrayResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Employee `json:"data"`
}

// Employee struct provides basic employee information
type Employee struct {
	FirstName string `gorm:"type:varchar(100)" json:"firstname"`
	LastName  string `gorm:"type:varchar(100)" json:"lastname"`
	Age       uint   `json:"age"`
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
		resp := ErrorResponse{
			Status:  http.StatusServiceUnavailable,
			Message: "failed",
		}
		c.JSON(http.StatusServiceUnavailable, resp)
	} else {
		resp := ArrayResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    emps,
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetOneEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp Employee
	err := db.Where("id = ?", id).First(&emp).Error

	if err != nil {
		resp := ErrorResponse{
			Status:  http.StatusServiceUnavailable,
			Message: "failed",
		}
		c.JSON(http.StatusServiceUnavailable, resp)
	} else {
		resp := Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    emp,
		}
		c.JSON(http.StatusOK, resp)
	}
}

func CreateEmployee(c *gin.Context) {
	var emp Employee
	err := c.BindJSON(&emp)

	if err != nil {
		resp := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Incorrect body",
		}
		c.JSON(http.StatusBadRequest, resp)
	} else {
		err := db.Create(&emp).Error
		if err != nil {
			resp := ErrorResponse{
				Status:  http.StatusServiceUnavailable,
				Message: "failed",
			}
			c.JSON(http.StatusServiceUnavailable, resp)
		} else {
			resp := Response{
				Status:  http.StatusOK,
				Message: "success",
				Data:    emp,
			}
			c.JSON(http.StatusOK, resp)
		}
	}
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp Employee
	if e := c.BindJSON(&emp); e != nil {
		resp := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Incorrect body",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := db.Where("id = ?", id).First(&emp).Error

	if err != nil {
		resp := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "failed",
		}
		c.JSON(http.StatusBadRequest, resp)
	} else {
		db.Save(&emp)
		resp := Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    emp,
		}
		c.JSON(http.StatusOK, resp)
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
		resp := Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    emp,
		}
		c.JSON(http.StatusOK, resp)
	}
}
