package controller

import (
	"booking-system/helper"
	"booking-system/model/web"
	"booking-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MidtransController struct {
	service service.MidtransService
	rg      *gin.RouterGroup
}

func (m *MidtransController) createHandler(ctx *gin.Context) {
	var request web.MidtransRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.PanicIfError(err)
	}

	midtransResponse := m.service.Create(ctx, request)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   midtransResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (m *MidtransController) webhookHandler(ctx *gin.Context) {
	var notificationPayload web.MidtransNotification
	if err := ctx.ShouldBindJSON(&notificationPayload); err != nil {
		helper.PanicIfError(err)
	}

	err := m.service.HandleWebhook(ctx, notificationPayload)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Webhook received successfully",
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (m *MidtransController) Router() {
	router := m.rg.Group("/midtrans")
	router.POST("/create", m.createHandler)
	router.POST("/webhook", m.webhookHandler)
}

func NewMidtransController(mS service.MidtransService, rg *gin.RouterGroup) *MidtransController {
	return &MidtransController{service: mS, rg: rg}
}
