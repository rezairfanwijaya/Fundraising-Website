package transaction

import user "github.com/rezairfanwijaya/Fundraising-Website/users"

// struct input untuk menampilkan transaksi berdasrkan campaign id
type GetTransactionsCampaignInput struct {
	CampaignId int `uri:"id" binding:"required"`
	User       user.User
}

// struct input untuk menyimpan input transaksi user
type CreateTransactionInput struct {
	Amount     int       `json:"amount"`
	CampaignId int       `json:"campaign_id"`
	User       user.User `json:"user"`
}
