package database

import (
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
	"fmt"
)

type Babe struct {
	gorm.Model
	ID uint
	Name string
	Gender string
	Age string
}

func Connect() (db *gorm.DB) {
	dsn := "homestead:secret@tcp(192.168.5.10:3306)/db_babeslist?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
    	panic("failed to connect database")
  	}
	fmt.Println("Connected to database!")
	db.AutoMigrate(&Babe{})

	return db
}