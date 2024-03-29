package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// buat interface service, yang berisi kontrak dari aktifitas user di website (register, login, etc)
type Service interface {
	RegisterUser(user RegisterUserInput) (User, error)
	Login(inputLogin LoginInput) (User, error)
	EmailIsAvaliable(inputEmail EmailInput) (User, error)
	UpdateAvatar(id int, filelocation string) (User, error)
	GetUserByID(id int) (User, error)
}

// buat internal struct untuk menampung repository, kita butuh repositoy agar bisa mengakses koneksi database dan juga function save data ke database
type service struct {
	repository Repository
}

// bikin new service untuk membuat struct internal service bisa digunakan
func NewService(repository Repository) *service {
	return &service{repository}
}

// function untuk register user
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	// cek apakah email sudah terdaftar atau belum
	userFromDB, err := s.repository.FindEmail(input.Email)
	if err != nil {
		return User{}, err
	}

	if userFromDB.Id != 0 {
		return User{}, errors.New("email telah terdaftar")
	}

	// proses maping struct input user ke struct representasi tabel user
	var user User
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email

	// encrypt password
	passENC, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passENC)
	user.Role = "user"

	// save ke database
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// function untuk login user
func (s *service) Login(inputLogin LoginInput) (User, error) {
	// kita ambil email dan password yang diinput user
	email := inputLogin.Email
	password := inputLogin.Password

	// cek email
	user, err := s.repository.FindEmail(email)
	if err != nil {
		return user, err
	}

	// jika email tidak ketemua (id=0)
	if user.Id == 0 {
		return user, errors.New("email tidak cocok")
	}

	// jika ditemukan maka cek passwordnya
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("password salah")
	}

	// jika lolos validasi email dan password
	return user, nil

}

// function untuk cek email
func (s *service) EmailIsAvaliable(inputEmail EmailInput) (User, error) {
	// ambil email
	email := inputEmail.Email

	// cari email
	user, err := s.repository.FindEmail(email)
	if err != nil {
		return user, err
	}

	// jika email sudah ada
	if user.Id != 0 {
		return user, errors.New("email sudah terdaftar")
	}

	return user, nil
}

// function untuk update avatar
func (s *service) UpdateAvatar(id int, filelocation string) (User, error) {
	// cari id
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, errors.New("user tidak ditemukan")
	}

	// update avatar
	user.AvatarFileName = filelocation

	// simpan data user hasil update
	newUser, err := s.repository.Update(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// function untuk mencari user berdasarkan id
func (s *service) GetUserByID(id int) (User, error) {
	// panggil repo findbyid
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	// cek id, jika 0 berarti ga ada user
	if user.Id == 0 {
		return user, errors.New("id tidak ditemukan")
	}

	return user, nil

}
