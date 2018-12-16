package controllers

import (
	"net/http"
	"simple_api_go_gorm/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// get all classes
func (idb *InDB) GetClasses(c *gin.Context) {
	var (
		classes []structs.Class
		result  gin.H
	)

	idb.DB.Find(&classes)
	if len(classes) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": classes,
			"count":  len(classes),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetClassAndStudent(c *gin.Context) {
	var (
		class  structs.Class
		result gin.H
	)
	id := c.Param("class_id")
	err := idb.DB.Where("id = ?", id).First(&class).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": class,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// insert class
func (idb *InDB) CreateClass(c *gin.Context) {
	var (
		class  structs.Class
		result gin.H
	)
	class_name := c.PostForm("class_name")
	class_time := time.Now()
	room := c.PostForm("room")

	class.Class_Name = class_name
	class.Class_Time = class_time
	class.Room = room
	idb.DB.Create(&class)
	result = gin.H{
		"result": class,
	}
	c.JSON(http.StatusOK, result)
}

// insert student
func (idb *InDB) CreateStudent(c *gin.Context) {
	var (
		student structs.Student
		result  gin.H
	)
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	class_id := c.PostForm("class_id")

	//i, err := strconv.ParseInt("1", 10, 16)

	id, err := strconv.Atoi(class_id)

	if err != nil {

	}

	student.Name = name
	student.Phone = phone
	student.Class_ID = id
	idb.DB.Create(&student)
	result = gin.H{
		"result": student,
	}
	c.JSON(http.StatusOK, result)
}
