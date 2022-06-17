package campaign

// bikin kontrak
type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
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
