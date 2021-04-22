package patient

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/patient"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetPDoctorSubscribedPatients(c echo.Context) ([]*model.Patient, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	result,err := op.GetDoctorPatients(claims.Sub)
	return result, err

}

