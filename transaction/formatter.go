package transaction

import "time"

// struct format response campaign transaction
type CamapaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// function untuk format single campaign transaction
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
