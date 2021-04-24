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

func GetAllNewBookingRequestsForDoctor(c echo.Context) ([]*model.BookingRequest, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	result,err := op.NewBookingRequestsForDoctor(claims.Sub)
	return result, err
}
