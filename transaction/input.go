package transaction

import user "github.com/rezairfanwijaya/Fundraising-Website/users"

// struct input untuk menampilkan transaksi berdasrkan campaign id
type GetTransactionsCampaignInput struct {
	CampaignId int `uri:"id" binding:"required"`
	User       user.User
}
