package transaction

import "time"

// struct format response campaign transaction by campaign id
type CamapaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// stuct format response campaign transactin by user id
type UserTransactionFormatter struct {
	Id int `json:"id"`
	Amount int `json:"amount"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Campaign CampaingFormatter `json:"campaign"` 
}

type CampaingFormatter struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

// function untuk format single campaign transaction transaction by campaign id
func FormatCampaignTransaction(transaction Transaction) CamapaignTransactionFormatter {
	campaignTransactionFormatter := CamapaignTransactionFormatter{}
	campaignTransactionFormatter.Id = transaction.Id
	campaignTransactionFormatter.Name = transaction.User.Name
	campaignTransactionFormatter.Amount = transaction.Amount
	campaignTransactionFormatter.CreatedAt = transaction.CreatedAt

	return campaignTransactionFormatter
}

// function untuk format many campaign transaction
func FormatCampaignTransactions(transactions []Transaction) []CamapaignTransactionFormatter {
	campaignTransactions := []CamapaignTransactionFormatter{}
	for _, transaction := range transactions {
		campaignTransaction := FormatCampaignTransaction(transaction)
		campaignTransactions = append(campaignTransactions, campaignTransaction)
	}

	return campaignTransactions
}

// function untuk format single campaing transaction by user id
func FormatUserTransaction(transactions Transaction) UserTransactionFormatter {
	userTransactionFormatter := UserTransactionFormatter{}
	userTransactionFormatter.Id = transactions.Id
	userTransactionFormatter.Amount = transactions.Amount
	userTransactionFormatter.Status = transactions.Status
	userTransactionFormatter.CreatedAt = transactions.CreatedAt

	campaign := CampaingFormatter{}
	campaign.Name = transactions.Campaign.Name
	campaign.ImageURL = ""
	if len(transactions.Campaign.CampaignImages) > 0 {
		campaign.ImageURL = transactions.Campaign.CampaignImages[0].FileName
	}

	userTransactionFormatter.Campaign = campaign

	return userTransactionFormatter
}

// function untuk format many campaign transaction by user id
 func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	// definisi return
	userTransactionsFormatter := []UserTransactionFormatter{}

	// cek apakah ada data 
	if len(transactions) == 0 {
		return userTransactionsFormatter
	}

	// lakukan format
	for _, transaction := range transactions {
		formatUserTransaction := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, formatUserTransaction)
	}

	// return
	return userTransactionsFormatter
 }