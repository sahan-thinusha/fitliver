package availability


import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/availability"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AddAvailability(c echo.Context) (*model.Availability, error) {
	availability := model.Availability{}
	if err := c.Bind(&availability); err != nil {
		return nil, err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.AddAvailabiility(&availability,claims.Sub)
	return result, err
}