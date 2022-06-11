package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/Fundraising-Website/handler"
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
	// handler user
	userHandler := handler.NewUserHandler(userService)

	// http server
	router := gin.Default()
	// api versioning
	api := router.Group("api/v1")
	// routing
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/session", userHandler.LoginUser)
	api.GET("/email", userHandler.CheckEmail)

	// run server
	router.Run(":7070")
}
