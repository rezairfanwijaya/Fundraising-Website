package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/Fundraising-Website/helper"
	"github.com/rezairfanwijaya/Fundraising-Website/transaction"
	user "github.com/rezairfanwijaya/Fundraising-Website/users"
)

// struct internal untuk menampung dependensi
type transactionHandler struct {
	service transaction.Service
}

// func new handler untuk dipakai di main
func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

// handler untuk get transaction berdasarkan campaign id
func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	// ambil uri berupa id campaign
	var input transaction.GetTransactionsCampaignInput

	// binding
	err := c.ShouldBindUri(&input)
	if err != nil {
		myErr := helper.ErrorFormater(err)
		data := gin.H{"error": myErr}
		response := helper.ResponsAPI("failed to get transaction", "failed", http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cek siapakah user yang melakukan request
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// panggil service
	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		data := gin.H{"error": err.Error()}
		response := helper.ResponsAPI("failed to get transaction", "failed", http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return
	response := helper.ResponsAPI("success to get transaction", "success", http.StatusOK, transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
