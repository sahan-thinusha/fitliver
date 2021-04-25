package payment

import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/payment"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func BookingPayment(c echo.Context) (*model.BookingPayment, error) {
	token := c.QueryParam("token")
	payment := model.BookingPayment{}
	if error := c.Bind(&payment); error != nil {
		return nil, error
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.ProceedBookingPayment(&payment,token,claims.Sub)
	return result, err
}


