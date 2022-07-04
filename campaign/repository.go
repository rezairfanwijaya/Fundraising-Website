package campaign

import (
	"gorm.io/gorm"
)

// definisi kontrak
type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userID int) ([]Campaign, error)
	FindById(id int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
}

// struct internal
type repository struct {
	db *gorm.DB
}

// func instance untuk dipakai dimain
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func untuk mencari semua campaign
func (r *repository) FindAll() ([]Campaign, error) {
	// slice penampung campaign
	var campaigns []Campaign

	// query
	// prealod akan mengambil tabel relasi yang dimiliki oleh tabel campaign yaitu tabel campaign_images
	// untuk params pada preload itu yang pertama berupa nama field si struct entity campaignimages dan yang kedua adalah kondisi yang kita inginkan
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	// return
	return campaigns, nil
}

// func untuk mencari campaign berdasarkan id user
func (r *repository) FindByUserId(userID int) ([]Campaign, error) {
	// slice penampung campaign
	var campaigns []Campaign

	// query
	// prealod akan mengambil tabel relasi yang dimiliki oleh tabel campaign yaitu tabel campaign_images
	// untuk params pada preload itu yang pertama berupa nama field si struct entity campaignimages dan yang kedua adalah kondisi yang kita inginkan
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// function untuk mencari campaign by id
func (r *repository) FindById(id int) (Campaign, error) {
	// deklarasi return
	var campaign Campaign

	// query
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", id).Find(&campaign).Error

	// error handling
	if err != nil {
		return campaign, err
	}

	// return
	return campaign, nil
}

// function untuk menyimpan data campaign
func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// function untuk melakukan update campaign
func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
