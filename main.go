package main

import (
	// "fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/Fundraising-Website/auth"
	// "github.com/rezairfanwijaya/Fundraising-Website/campaign"
	"github.com/rezairfanwijaya/Fundraising-Website/handler"
	"github.com/rezairfanwijaya/Fundraising-Website/helper"
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
	// service auth
	authService := auth.NewServiceAuth()
	// handler user
	userHandler := handler.NewUserHandler(userService, authService)

	// campaignRepo := campaign.NewRepository(db)
	// camp, err := campaignRepo.FindAll()
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }

	// for _, v := range camp {
	// 	if len(v.CampaignImages) > 0 {
	// 		fmt.Println(v.Name, v.CampaignImages[0].FileName)
	// 	} else {
	// 		fmt.Println(v.Name)
	// 	}
	// }

	// id, err := campaignRepo.FindByUserId(1)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }

	// for _, v := range id {
	// 	fmt.Println(v.Description)
	// 	if len(v.CampaignImages) > 0 {
	// 		fmt.Println(v.CampaignImages[0].FileName)
	// 	}
	// }

	// return

	// http server
	router := gin.Default()
	// api versioning
	api := router.Group("api/v1")
	// routing
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/session", userHandler.LoginUser)
	api.POST("/email", userHandler.CheckEmail)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UpdateAvatar)

	// run server
	router.Run(":7070")
}

// function middleware auth
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil nilai header yang sudah kita set namanya Authorization
		authHeader := c.GetHeader("Authorization")

		// cek apakah nilai authorization memiliki Bearer
		// karena nanti kita akan set nilai token seperti ini "Bearer djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"
		if !strings.Contains(authHeader, "Bearer") {
			respons := helper.ResponsAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// lalu kita pisahkan tokennya berdasarkan spasi
		// before "Bearer djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"
		// after ["Bearer"] ["djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"]
		tokenString := ""
		arraytoken := strings.Split(authHeader, " ")
		if len(arraytoken) == 2 {
			tokenString = arraytoken[1]
		}

		// validasi token
		token, err := authService.ValidasiToken(tokenString)
		if err != nil {
			respons := helper.ResponsAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// ambil data dalam token
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			respons := helper.ResponsAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// data id user diambil
		userId := int(claim["user_id"].(float64))

		// data user diambil berdasarkan id
		user, err := userService.GetUserByID(userId)
		if err != nil {
			respons := helper.ResponsAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// set context
		c.Set("currentUser", user)

	}
}
