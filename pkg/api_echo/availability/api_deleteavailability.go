package availability

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/availability"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func DeleteAvailability(c echo.Context) (*model.Availability, error) {
	id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	availabilityID,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.DeleteAvailability(availabilityID,claims.Sub)
	return result, err
}
