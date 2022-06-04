package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/Fundraising-Website/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect ke database
	r := gin.Default()

	r.GET("/", ShowAll)

	r.Run(":7070")

}

func ShowAll(c *gin.Context) {
	dsn := "root@tcp(127.0.0.2:3306)/fundraishing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var user []models.User

	db.Find(&user)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"code": http.StatusOK,
	})
}
