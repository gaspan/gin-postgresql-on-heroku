package employee

import (
	"employee-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var employees = []model.Employee{}

// List all employees
func List(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

// Add a new employee
func Add(c *gin.Context) {
	var employee model.Employee
	err := c.BindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	employee.ID = uint(len(employees))
	employees = append(employees, employee)

	resp := model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    employee,
	}

	c.JSON(http.StatusOK, resp)
}

// Update an existing employee
func Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employee model.Employee
	err := c.BindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i, e := range employees {
		if e.ID == uint(id) {
			employee.ID = e.ID
			employees[i] = employee
		}
	}

	resp := model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    employee,
	}

	c.JSON(http.StatusOK, resp)
}

// Delete an employee
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, e := range employees {
		if e.ID == uint(id) {
			employees = append(employees[:i], employees[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, employees)
}
