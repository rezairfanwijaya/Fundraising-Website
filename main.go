package main

import (
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

	// repo user
	userRepo := user.NewRepository(db)
	// service user
	userService := user.NewService(userRepo)

	// input register user
	var newUser user.RegisterUserInput
	newUser.Name = "Reza Irfan Abdas"
	newUser.Occupation = "Mahasiswa"
	newUser.Email = "rezairfanabdas@gmail.com"
	newUser.Password = "123"
	userService.RegisterUser(newUser)
}
