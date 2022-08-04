package transaction

import user "github.com/rezairfanwijaya/Fundraising-Website/users"

// struct input untuk menampilkan transaksi berdasrkan campaign id
type GetTransactionsCampaignInput struct {
	CampaignId int `uri:"id" binding:"required"`
	User       user.User
}

// struct input untuk menyimpan input transaksi user
type CreateTransactionInput struct {
	Amount     int       `json:"amount" binding:"required"`
	CampaignId int       `json:"campaign_id" binding:"required"`
	User       user.User `json:"user" binding:"required"`
}
