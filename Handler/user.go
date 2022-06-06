package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/rezairfanwijaya/Fundraising-Website/helper"
	user "github.com/rezairfanwijaya/Fundraising-Website/users"
)

// bikin struct internal userHandler yang akan menyimpan service dari user yang di dalam nya berisi akses save data ke database

type userHandler struct {
	userService user.Service
}

// bikin NewUserHandler untuk membuat userHandler berfungsi di main
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// bikin handler register
func (u *userHandler) RegisterUser(c *gin.Context) {
	// variable menampung inputan register user
	var inputUser user.RegisterUserInput

	// binding
	err := c.ShouldBindJSON(&inputUser)
	if err != nil {
		respons := helper.ResponsAPI("Gagal menyimpan data", "Gagal", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, respons)

		return
	}

	// save ke database
	newUser, err := u.userService.RegisterUser(inputUser)
	if err != nil {
		respons := helper.ResponsAPI("Gagal menyimpan data", "Gagal", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, respons)

		return
	}

	// sebelum di tampilkan data ke user maka data User dari inputan user harus kita formating dulu sesuai yg diminta pada helper
	userFormat := user.UserFormatter(newUser, "tokenuser")

	respons := helper.ResponsAPI("Berhasil menyimpan data", "sukses", http.StatusBadRequest, userFormat)

	c.JSON(http.StatusOK, respons)
}
