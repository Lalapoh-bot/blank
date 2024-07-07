package service

import (
	"booking-system/helper"
	"booking-system/model/web"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	Create(c *gin.Context, request web.MidtransRequest) web.MidtransResponse
	HandleWebhook(ctx context.Context, notification web.MidtransNotification) error
}

type MidtransServiceImpl struct {
	Validate *validator.Validate
}

func NewMidtransServiceImpl(validate *validator.Validate) *MidtransServiceImpl {
	return &MidtransServiceImpl{
		Validate: validate,
	}
}

func (service *MidtransServiceImpl) Create(c *gin.Context, request web.MidtransRequest) web.MidtransResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper.PanicIfError(err)
	}

	var snapClient = snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	user_id := strconv.Itoa(request.UserId)

	custAddress := &midtrans.CustomerAddress{
		FName:       "John",
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-User-" + user_id + "-" + request.ItemID,
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "dibsi",
			LName:    "Teletubis",
			Email:    "dibsi@teletubbies.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + request.ItemID,
				Qty:   1,
				Price: request.Amount,
				Name:  request.ItemName,
			},
		},
	}

	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		helper.PanicIfError(errSnap.GetRawError())
	}

	midtransReponse := web.MidtransResponse{
		Token:       response.Token,
		RedirectUrl: response.RedirectURL,
	}

	return midtransReponse
}

func (service *MidtransServiceImpl) HandleWebhook(ctx context.Context, notification web.MidtransNotification) error {
	fmt.Println("Received webhook notification:", notification)

	switch notification.TransactionStatus {
	case "capture":
		if notification.FraudStatus == "accept" {
			// Update your transaction status in the database to 'success'
			fmt.Println("Transaction status: capture and accepted")
		} else {
			// Handle other fraud status
			fmt.Println("Transaction status: capture but fraud status not accepted")
		}
	case "settlement":
		// Update your transaction status in the database to 'success'
		fmt.Println("Transaction status: settlement")
	case "deny", "expire", "cancel":
		// Update your transaction status in the database to 'failure'
		fmt.Println("Transaction status: ", notification.TransactionStatus)
	default:
		return errors.New("unknown transaction status")
	}

	return nil
}
