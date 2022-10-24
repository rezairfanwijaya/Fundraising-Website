package payment

import (
	"log"
	"strconv"

	"github.com/rezairfanwijaya/Fundraising-Website/helper"
	user "github.com/rezairfanwijaya/Fundraising-Website/users"
	"github.com/veritrans/go-midtrans"
)

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	// get key from env file
	env, err := helper.GetENV("./.env")
	if err != nil {
		log.Fatal(err)
	}
	ServerKey := env["SERVER_KEY"]
	ClientKey := env["CLIENT_KEY"]

	// setup untuk midtrans
	midclient := midtrans.NewClient()
	midclient.ServerKey = ServerKey
	midclient.ClientKey = ClientKey
	midclient.APIEnvType = midtrans.Sandbox

	// deklarasi gateway
	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	// siapkan payload untuk get redirect url
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
	}

	// get redirect url
	snapResponse, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapResponse.RedirectURL, nil
}
