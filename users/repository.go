package user

import "gorm.io/gorm"

// bikin interface untuk kontrak
type Repository interface {
	Save(user User) (User, error)
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
