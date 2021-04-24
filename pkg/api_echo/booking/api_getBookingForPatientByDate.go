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

func GetAllBookingsForPatientByDate(c echo.Context) ([]*model.Patient_Booking, error) {
	date := c.QueryParam("date")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	result,err := op.BookingsForDoctorByDate(date,claims.Sub)
	return result, err
}
