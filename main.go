package main

import (
	"fmt"
	"log"

	user "github.com/rezairfanwijaya/Fundraising-Website/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect ke database
	dsn := "root:@tcp(127.0.0.1:3306)/fundraishing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepository(db)
	newUser := user.User{
		Id:   3,
		Name: "skuteowro",
	}

	res, err := userRepo.Save(newUser)
	if err != nil {
		log.Printf("Failed Save: %v", err)
	}

	fmt.Println(res)
}
