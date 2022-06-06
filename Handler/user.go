package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		// bikin variable untuk menampung error
		var errors []string

		// ambil error
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		// bikin format json baru untuk masuk ke data di responsAPI
		myErr := gin.H{"error": errors}

		// masukan error ke responsAPI
		respons := helper.ResponsAPI("Gagal menyimpan data", "Gagal", http.StatusUnprocessableEntity, myErr)
		c.JSON(http.StatusUnprocessableEntity, respons)

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

	respons := helper.ResponsAPI("Berhasil menyimpan data", "sukses", http.StatusOK, userFormat)

	c.JSON(http.StatusOK, respons)
}
