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

		// panggil function di helper untuk mengolah error
		errors := helper.ErrorFormater(err)

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
		respons := helper.ResponsAPI("Gagal menyimpan data", "Gagal", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, respons)

		return
	}

	// sebelum di tampilkan data ke user maka data User dari inputan user harus kita formating dulu sesuai yg diminta pada helper
	userFormat := user.UserFormatter(newUser, "tokenuser")

	respons := helper.ResponsAPI("Berhasil menyimpan data", "sukses", http.StatusOK, userFormat)

	c.JSON(http.StatusOK, respons)
}

// bikin handler login
func (u *userHandler) LoginUser(c *gin.Context) {

	// definisikan struct input login
	var input user.LoginInput

	// binding
	err := c.ShouldBindJSON(&input)
	if err != nil {

		// panggil error formater
		myError := helper.ErrorFormater(err)

		// masukan ke template respons API
		response := helper.ResponsAPI("Login gagal", "Gagal", http.StatusUnprocessableEntity, myError)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cek email dan password
	loggedUser, err := u.userService.Login(input)
	if err != nil {
		response := helper.ResponsAPI("Login gagal", "Gagal", http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// jika email dan password cocok maka user harus diformat terlebih dahulu
	formatUser := user.UserFormatter(loggedUser, "tokentokentoken")

	// lalu masukan ke respons
	response := helper.ResponsAPI("Login berhasil", "Sukses", http.StatusOK, formatUser)
	c.JSON(http.StatusUnprocessableEntity, response)
}
