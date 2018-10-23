package main

import (
	"employee-api/handler"
	"fmt"
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
		"host="+os.Getenv("DB_HOST")+" user="+os.Getenv("DB_USER")+
			" dbname="+os.Getenv("DB_NAME")+" sslmode=disable password="+
			os.Getenv("DB_PASS"))

	fmt.Print("host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") + " sslmode=disable password=" +
		os.Getenv("DB_PASS"))

	if err != nil {
		fmt.Print(err)
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
