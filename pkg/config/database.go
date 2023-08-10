package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connector *gorm.DB

func Connect() {
	dsn := "user:pass@/dbname?charset=utf8&parseTime=True&loc=Local" //TODO: add func to make custom dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	connector = db
}

func GetDB() *gorm.DB {
	return connector
}
