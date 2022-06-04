package main

import (
	"fmt"
	"log"

	"github.com/rezairfanwijaya/Fundraising-Website/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect ke database
	dsn := "root@tcp(127.0.0.2:3306)/fundraishing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println(db)

	var user []models.User

	db.Find(&user)

	for _, v := range user {
		fmt.Println(v)
	}
}
