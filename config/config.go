package config

import (
	"simple_api_go_gorm/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/golang_api_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Class{}, structs.Student{})
	return db
}
