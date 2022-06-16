package campaign

import "gorm.io/gorm"

// definisi kontrak
type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(user_id int) ([]Campaign, error)
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
func (r *repository) FindByUserId(user_id int) ([]Campaign, error) {
	// slice penampung campaign
	var campaigns []Campaign

	// query
	// prealod akan mengambil tabel relasi yang dimiliki oleh tabel campaign yaitu tabel campaign_images
	// untuk params pada preload itu yang pertama berupa nama field si struct entity campaignimages dan yang kedua adalah kondisi yang kita inginkan
	err := r.db.Where("user_id = ?", user_id).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}