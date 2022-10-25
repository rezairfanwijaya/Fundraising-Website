package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/Fundraising-Website/auth"
	helper "github.com/rezairfanwijaya/Fundraising-Website/helper"
	user "github.com/rezairfanwijaya/Fundraising-Website/users"
)

// bikin struct internal userHandler yang akan menyimpan service dari user yang di dalam nya berisi akses save data ke database

type userHandler struct {
	userService  user.Service
	tokenservice auth.Service
}

// bikin NewUserHandler untuk membuat userHandler berfungsi di main
func NewUserHandler(userService user.Service, tokenService auth.Service) *userHandler {
	return &userHandler{userService, tokenService}
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

	// generate token
	token, err := u.tokenservice.GenerateToken(newUser.Id)
	if err != nil {
		respons := helper.ResponsAPI("Gagal membuat token", "Gagal", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, respons)
	}

	// sebelum di tampilkan data ke user maka data User dari inputan user harus kita formating dulu sesuai yg diminta pada helper
	userFormat := user.UserFormatter(newUser, token)

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

	// generate token
	token, err := u.tokenservice.GenerateToken(loggedUser.Id)
	if err != nil {
		respons := helper.ResponsAPI("Gagal membuat toke", "Gagal", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// jika email dan password cocok maka user harus diformat terlebih dahulu
	formatUser := user.UserFormatter(loggedUser, token)

	// lalu masukan ke respons
	response := helper.ResponsAPI("Login berhasil", "Sukses", http.StatusOK, formatUser)
	c.JSON(http.StatusOK, response)
}

// handler check email
func (u *userHandler) CheckEmail(c *gin.Context) {
	// inisiasi email input
	var email user.EmailInput

	// binding
	err := c.ShouldBindJSON(&email)
	if err != nil {
		myErr := helper.ErrorFormater(err)
		respons := helper.ResponsAPI("Gagal cek email", "Gagal", http.StatusUnprocessableEntity, myErr)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	// cek email
	_, err = u.userService.EmailIsAvaliable(email)
	if err != nil {
		respons := helper.ResponsAPI("Email Telah Terpakai", "Gagal", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	data := gin.H{
		"message": "Email diperbolehkan",
	}
	respons := helper.ResponsAPI("Email Diterima", "Sukses", http.StatusOK, data)
	c.JSON(http.StatusOK, respons)
}

// handler update avatar
func (u *userHandler) UpdateAvatar(c *gin.Context) {
	// kita harus tangkap file gambar yang diupload oleh user
	file, err := c.FormFile("avatar") // string avatar harus sama dengan atribut name di tag input html
	if err != nil {
		respons := helper.ResponsAPI("Avatar gagal diupload", "Gagal", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// kita harus menyimpan file gambar di local kita
	// file.Filename adalah nama gambar yang diupload user (john.png)
	path := "images/" + file.Filename

	// proses simpan gambar ka local
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		respons := helper.ResponsAPI("Avatar gagal diupload", "Gagal", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, respons)
		return
	}

	// proses update file
	// kita tentuin siapa user yang akan mengupate avatar, ini didapat melalui id yang sudah disimpan didalam context ketika melwati middleware. Jadi kita tinggal ambil context nya saja
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.Id

	_, err = u.userService.UpdateAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		respons := helper.ResponsAPI("Avatar gagal diupload", "Gagal", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, respons)
		return
	}
	data := gin.H{
		"is_uploaded": true,
	}

	respons := helper.ResponsAPI("Avatar berhasil diupload", "Sukses", http.StatusOK, data)
	c.JSON(http.StatusOK, respons)

}

// handle untuk mendapatkan user yang sedang login
func (u *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userFormat := user.UserFormatter(currentUser, "")
	response := helper.ResponsAPI("success get current user", "success", http.StatusOK, userFormat)
	c.JSON(http.StatusOK, response)
}
