package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Class struct {
	Class_Name string
	Room       string
	Class_Time time.Time `gorm:"type:time"`
	gorm.Model
	Student []Student
}

type Student struct {
	gorm.Model
	Name     string
	Phone    string
	Class_ID int
}
