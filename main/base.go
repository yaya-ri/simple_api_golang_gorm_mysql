package main

import (
	"simple_api_go_gorm/config"
	"simple_api_go_gorm/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/classes", inDB.GetClasses)
	router.GET("/students/:class_id", inDB.GetClassAndStudent)
	router.POST("/class", inDB.CreateClass)
	router.POST("/student", inDB.CreateStudent)

	router.Run(":3000")
}
