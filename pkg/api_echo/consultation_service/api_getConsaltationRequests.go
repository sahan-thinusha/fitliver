package consultation_service

import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/consultation_service"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetAllNewConsultationRequestsForDoctor(c echo.Context) ([]*model.ConsultationRequest, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	result,err := op.NewConsultationRequestsForDoctor(claims.Sub)
	return result, err
}
