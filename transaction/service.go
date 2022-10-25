package transaction

import (
	"errors"
	"strconv"

	"github.com/rezairfanwijaya/Fundraising-Website/campaign"
	"github.com/rezairfanwijaya/Fundraising-Website/helper"
	"github.com/rezairfanwijaya/Fundraising-Website/payment"
)

// bikin kontrak
type Service interface {
	GetTransactionByCampaignId(input GetTransactionsCampaignInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProsesPayment(input MidtransNotifications) error
}

// bikin internal struct untuk meletakan dependensi
type service struct {
	repository     Repository
	campaignRepo   campaign.Repository
	paymentService payment.Service
}

// bikin new service untuk dipakai di main
func NewService(
	repository Repository,
	campaignRepo campaign.Repository,
	paymentService payment.Service,

) *service {
	return &service{repository, campaignRepo, paymentService}
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

// function untuk menyimpan data transaksi user
func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	// asign value
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignId
	transaction.Amount = input.Amount
	transaction.UserId = input.User.Id
	transaction.Status = "pending"

	// save data
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, errors.New("failed to save transaction")
	}

	// set payment transaction to get payment url
	paymentTransaction := payment.Transaction{
		Id:     newTransaction.Id,
		Amount: newTransaction.Amount,
	}
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	// update paymentURL to table
	newTransaction.PaymentURL = paymentURL
	code := helper.GenerateCodeTransaction(newTransaction.Id)
	newTransaction.Code = code
	transactionWithPaymentURL, err := s.repository.Update(newTransaction)
	if err != nil {
		return transactionWithPaymentURL, err
	}

	// return
	return newTransaction, nil
}

// function untuk handle notifikasi pembayaran dari midtrans ketika user berhasil melakukan pembayaran
func (s *service) ProsesPayment(input MidtransNotifications) error {
	// get transaction by id from midtrans notification
	transaction_id, _ := strconv.Atoi(input.OrderId)
	transaction, err := s.repository.GetById(transaction_id)
	if err != nil {
		return err
	}

	// handle status
	if input.PaymentType == "credit_card" &&
		input.TransactionStatus == "capture" &&
		input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" ||
		input.TransactionStatus == "expire" ||
		input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	// update transaction with new status
	transactionUpdated, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	// find campaign to update that
	campaign, err := s.campaignRepo.FindById(transactionUpdated.CampaignID)
	if err != nil {
		return err
	}

	// cek status transaction from transactionUpdated
	// and update campaign
	if transactionUpdated.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + transactionUpdated.Amount
		_, err := s.campaignRepo.Update(campaign)
		if err != nil {
			return err
		}
	}

	// success
	return nil
}
