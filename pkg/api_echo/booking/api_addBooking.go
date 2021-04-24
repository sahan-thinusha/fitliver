package booking


import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/booking"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateBookingService(c echo.Context) (*model.BookingService, error) {
	bookingService := model.BookingService{}
	if error := c.Bind(&bookingService); error != nil {
		return nil, error
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.BookingServiceCreate(&bookingService,claims.Sub)
	return result, err
}
