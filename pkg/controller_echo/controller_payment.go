package controller_echo

import (
	"fitliver/pkg/api_echo/payment"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreatePayment(c echo.Context) error {
	result, err := payment.MakePayment(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func CreateBookingPayment(c echo.Context) error {
	result, err := payment.BookingPayment(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func APIControllerPayment(g *echo.Group) {
	g.POST("api/pay", CreatePayment)
	g.POST("api/booking/pay", CreateBookingPayment)

}
