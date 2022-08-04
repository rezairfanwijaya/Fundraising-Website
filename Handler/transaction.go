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

// handler untuk mengambil data transaksi berdasrkan userid
func (h *transactionHandler) GetTransactionByUserId(c *gin.Context) {
	// kita ambil data user yang melakukan request melalui JWT
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id

	// panggil service
	transactions, err := h.service.GetTransactionByUserId(userId)
	if err != nil {
		data := gin.H{"error": err.Error()}
		response := helper.ResponsAPI("failed to get transaction", "failed", http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return
	response := helper.ResponsAPI("success to get transaction", "success", http.StatusOK, transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)

}

// handler untuk membuat transaksi
func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	// tampung input user
	var input transaction.CreateTransactionInput

	// binding
	err := c.ShouldBindJSON(&input)
	if err != nil {
		myErr := helper.ErrorFormater(err)
		response := helper.ResponsAPI("failed to create transaction", "failed", http.StatusUnprocessableEntity, myErr)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// isi user ke variable input
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// panggil service untuk membuat data transaksi baru
	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		response := helper.ResponsAPI("failed to create new transaction", "failed", http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// response sukses
	response := helper.ResponsAPI("success create new transaction", "success", http.StatusOK, newTransaction)
	c.JSON(http.StatusOK, response)
}
