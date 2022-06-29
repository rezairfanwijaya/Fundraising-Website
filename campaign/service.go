package campaign

import "errors"

// bikin kontrak
type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignById(input InputCampaignDetail) (Campaign, error)
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
		return campaign, errors.New("gagal mengambil data campaign")
	}

	// return
	return campaign, nil
}
