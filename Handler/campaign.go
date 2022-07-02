package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/Fundraising-Website/campaign"
	"github.com/rezairfanwijaya/Fundraising-Website/helper"
	user "github.com/rezairfanwijaya/Fundraising-Website/users"
)

// bikin struct
type campaignHandler struct {
	service campaign.Service
}

// bikin function new handler
func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// handler getCampaign
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// nanti bentuk url akan seperti ini
	// v1/api/campaigns/:user_id
	// v1/api/campaigns?user_id=10
	// bearti kita harus tangkap params dari endpoint nya
	userId, _ := strconv.Atoi(c.Query("user_id"))

	// panggil function GetCampaigns
	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		myErr := helper.ErrorFormater(err)
		response := helper.ResponsAPI("Error to get campaigns", "error", http.StatusBadRequest, myErr)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponsAPI("List of campaigns", "success", http.StatusOK, campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

// handler get campaign by id
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// output endpoint yang diharapkan ---> /api/campaign/1
	// ambil input user
	var input campaign.InputCampaignDetail

	// binding via uri
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ResponsAPI("Get data failed", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// panggil service
	campaignDetail, err := h.service.GetCampaignById(input)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		response := helper.ResponsAPI("Get data failed", "Failed", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// formatter
	data := campaign.FormatDetailCampaign(campaignDetail)

	response := helper.ResponsAPI("Get data succes", "succes", http.StatusOK, data)
	c.JSON(http.StatusOK, response)

}

// handler untuk create campagin
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	// definisi input
	var input campaign.CreateCampaignInput

	// binding
	err := c.ShouldBindJSON(&input)
	if err != nil {
		myErr := helper.ErrorFormater(err)
		response := helper.ResponsAPI("failed to create campaign", "failed", http.StatusUnprocessableEntity, myErr)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// ambil data user dari jwt yg sudah disimpan di context
	curretUser := c.MustGet("currentUser").(user.User)

	// assign curretUser ke user yang ada diinput
	input.User = curretUser

	// panggil service
	data, err := h.service.CreateCampaign(input)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		response := helper.ResponsAPI("failed to create campaign", "failed", http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// format campaign
	newCampaign := campaign.FormatCampaign(data)
	response := helper.ResponsAPI("success to create campaign", "success", http.StatusOK, newCampaign)
	c.JSON(http.StatusOK, response)
}
