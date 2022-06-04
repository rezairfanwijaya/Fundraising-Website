package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/rezairfanwijaya/Fundraising-Website/User"
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

	var users []models.User

	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"code": http.StatusOK,
	})
}
