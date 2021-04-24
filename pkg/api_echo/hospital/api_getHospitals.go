package hospital

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/hospital"
	"github.com/labstack/echo/v4"
)

func GetHospitals(c echo.Context) ([]*model.Hospital, error) {
	result,err := op.GetHospitals()
	return result, err
}