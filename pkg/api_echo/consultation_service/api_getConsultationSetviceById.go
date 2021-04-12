package consultation_service

import (
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/consultation_service"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetConsultationServicesById(c echo.Context) (*model.ConsultationService, error) {
	id := c.Param("id")
	consultationServiceId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.ConsultationServiceById(consultationServiceId)
	return result, err
}
