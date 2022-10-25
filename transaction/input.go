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

// struct untuk handle notifikasi yang dikirim oleh midtrans
// https://docs.midtrans.com/en/after-payment/http-notification
type MidtransNotifications struct {
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	OrderId           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
}
