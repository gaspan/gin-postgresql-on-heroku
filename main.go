package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// An Employee struct represents data about an employee
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}

// An Response struct represents http response to the client from the server
type Response struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Employee `json:"data"`
}

var employees = []Employee{
	Employee{1, "Wasuwat", "Limsuparhat", 22},
	Employee{2, "Suepsakun", "Aiamlaoo", 22},
	Employee{3, "Sitthipon", "Songsaen", 23},
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	api.GET("employee", listEmployeeHandler)
	api.POST("employee", addEmployeeHandler)
	api.PUT("employee/:id", updateEmployeeHandler)
	api.DELETE("employee/:id", deleteEmployeeHandler)

	router.Run("localhost:8000")
}

func listEmployeeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

func addEmployeeHandler(c *gin.Context) {
	var employee Employee
	err := c.BindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	employee.ID = len(employees) + 1
	employees = append(employees, employee)

	resp := Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    employee,
	}

	c.JSON(http.StatusOK, resp)
}

func updateEmployeeHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employee Employee
	err := c.BindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i, e := range employees {
		if e.ID == id {
			employee.ID = e.ID
			employees[i] = employee
		}
	}

	resp := Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    employee,
	}

	c.JSON(http.StatusOK, resp)
}

func deleteEmployeeHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, employees)
}
