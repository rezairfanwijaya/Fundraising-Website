package transaction

import "gorm.io/gorm"

// bikin kontrak
type Repository interface {
	GetByCampaignId(campaignId int) ([]Transaction, error)
	GetByUserId(userId int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
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

// function untuk mencari transaksi untuk user tertentu
func (r *repository) GetByUserId(userId int) ([]Transaction, error) {
	// definisi return
	var transactions []Transaction

	// query
	err := r.db.Preload("Campaign.CampaignImages", "is_primary = 1").Where("user_id = ? ", userId).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	// return
	return transactions, nil
}

// function untuk menyimpan data transaksi user
func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

// function update
func (r *repository) Update(transaction Transaction) (Transaction, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}
