package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var connector *gorm.DB

type Config struct {
	User     string
	Password string
	DB       string
	Server   string
}

var DBConfig = Config{
	User:     "root",
	Password: "password",
	DB:       "gameshop",
	Server:   "localhost:8080",
}

func GetDSN(config *Config) string {
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.DB)
}

func Connect() {
	dsn := GetDSN(&DBConfig)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db: %w", err)
	}
	connector = db
	log.Println("Successful connection to db")
}

func GetDB() *gorm.DB {
	return connector
}
