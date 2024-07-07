package controller

import (
	"booking-system/service"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	service service.BookingService
	rg      *gin.RouterGroup
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {

}

func (c *BookingController) Router() {
	router := c.rg.Group("/court")
	router.POST("/payments", c.CreateBooking)

}
func NewBookingController(bookingService service.BookingService, rg *gin.RouterGroup) *BookingController {
	return &BookingController{service: bookingService, rg: rg}
}
