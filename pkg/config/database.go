package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connector *gorm.DB

func Connect() {
	dsn := "root:password@/gameshop?charset=utf8&parseTime=True&loc=Local" //TODO: add func to make custom dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	connector = db
	fmt.Println("connected")
}

func GetDB() *gorm.DB {
	return connector
}
