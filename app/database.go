package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:rakamin@tcp(127.0.0.1:3306)/rakamin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Connection to database established")
	return db, nil
}