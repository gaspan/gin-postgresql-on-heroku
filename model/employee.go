package model

import (
	"github.com/jinzhu/gorm"
)

// An Employee struct represents data about an employee
type Employee struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100)"`
	LastName  string `gorm:"type:varchar(100)"`
	Age       uint
}
