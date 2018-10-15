package model

// An Employee struct represents data about an employee
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}
