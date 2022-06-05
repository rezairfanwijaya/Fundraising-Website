package user

import "golang.org/x/crypto/bcrypt"

// buat interface service, yang berisi kontrak dari aktifitas user di website (register, login, etc)
type Service interface {
	RegisterUser(user RegisterUserInput) (User, error)
}

// buat internal struct untuk menampung repository, kita butuh repositoy agar bisa mengakses koneksi database dan juga function save data ke database
type service struct {
	repository Repository
}

// bikin new service untuk membuat struct internal service bisa digunakan
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

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
