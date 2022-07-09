package transaction

import "gorm.io/gorm"

// bikin kontrak
type Repository interface {
	GetByCampaignId(campaignId int) ([]Transaction, error)
}

// bikin internal struct
type repository struct {
	db *gorm.DB
}

// bikin function new repo untuk dipakai di main
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// function untuk mencari transaksi pada campaign tertentu
func (r *repository) GetByCampaignId(campaignId int) ([]Transaction, error) {
	// definisi return
	var transaction []Transaction

	// query
	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
