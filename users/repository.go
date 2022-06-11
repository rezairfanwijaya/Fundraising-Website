// repository ini menampung semua function yang berhubungan dengan database
package user

import "gorm.io/gorm"

// bikin interface untuk kontrak
type Repository interface {
	Save(user User) (User, error)
	FindEmail(email string) (User, error)
	FindById(id int) (User, error)
	Update(user User) (User, error)
}

// struct respository untuk menampung koneksi database
type repository struct {
	db *gorm.DB
}

// newrepo untuk mengisi koneksi ke struct repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// function untuk save data user ke database
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// function untuk cek email user
func (r *repository) FindEmail(email string) (User, error) {
	// inisiasi struct user sebagai nilai balikan
	var user User
	// proses pencarian email
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// function untuk mencari data user berdasarkan ID
func (r *repository) FindById(id int) (User, error) {
	// inisiasi struct user sebagai nilai balikan
	var user User

	// cari id di database
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// function untuk update user
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
