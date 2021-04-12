package consultation_service

import (
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/consultation_service"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetAllConsultationServices(c echo.Context) ([]*model.ConsultationService, error) {
	result,err := op.ConsultationServicesGetAll()
	return result, err
}
