package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

// bikin kontrak
type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignById(input InputCampaignDetail) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

// bikin struct untuk menampung repository (dependensi)
type service struct {
	repository Repository
}

// bikin newservice untuk dipanggil oleh public
func NewService(repository Repository) *service {
	return &service{repository}
}

// bikin function get campaign untuk mengembalikan semua campaign
func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	// cek apakah userID ada / nggak
	// jika ada berarti panggil function di repo yang findbyid
	// kalo ga ada berarti panggil function di repo yang findall
	if userID != 0 {
		campaigns, err := s.repository.FindByUserId(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// bikin function untuk mengambil data campaign berdasarakan id yang dikirim lewat endpoint
func (s *service) GetCampaignById(input InputCampaignDetail) (Campaign, error) {
	// panggil repo
	campaign, err := s.repository.FindById(input.Id)

	// error handling
	if err != nil {
		return campaign, errors.New("failed to get data campaign")
	}

	// return
	return campaign, nil
}

// function untuk membuat campaign
func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	// inisiasi campaign
	var campaign Campaign
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserId = input.User.Id

	// pembuatan slug
	// slug harus unik, maka kita masukan user id dinama slug nya
	// contoh : nama-campagin-10
	slugCandidate := fmt.Sprintf("%v %v", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugCandidate)

	// panggil repo
	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, errors.New("failed to save new campaign")
	}

	// return
	return newCampaign, nil
}
