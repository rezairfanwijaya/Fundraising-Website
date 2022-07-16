package transaction

import (
	"errors"

	"github.com/rezairfanwijaya/Fundraising-Website/campaign"
)

// bikin kontrak
type Service interface {
	GetTransactionByCampaignId(input GetTransactionsCampaignInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
}

// bikin internal struct untuk meletakan dependensi
type service struct {
	repository   Repository
	campaignRepo campaign.Repository
}

// bikin new service untuk dipakai di main
func NewService(repository Repository, campaignRepo campaign.Repository) *service {
	return &service{repository, campaignRepo}
}

// function untuk mengambil data transaksi berdasarkan campaign id
func (s *service) GetTransactionByCampaignId(input GetTransactionsCampaignInput) ([]Transaction, error) {
	// ambil campaign berdasarkan id
	campaign, err := s.campaignRepo.FindById(input.CampaignId)
	if err != nil {
		return []Transaction{}, errors.New("campaign not found")
	}

	// cek apakah yang melakukan request adalah user pemilik campaign
	if input.User.Id != campaign.UserId {
		return []Transaction{}, errors.New("access denied, only owner campaign can access")
	}

	// panggil repo
	transactions, err := s.repository.GetByCampaignId(input.CampaignId)
	if err != nil {
		return transactions, errors.New("transaction not found")
	}

	return transactions, nil
}

// function untuk mengambil data transaksi berdasarkan userid
func (s *service) GetTransactionByUserId(userId int) ([]Transaction, error) {
	// panggil repo
	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transactions, errors.New("user id not found")
	}

	// return
	return transactions, nil
}
