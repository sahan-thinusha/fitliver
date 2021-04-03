package doctor

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/doctor"
	"github.com/labstack/echo/v4"
)

func GetDoctors(c echo.Context) ([]*model.Doctor, error) {
	result,err := op.GetDoctors()
	return result, err

}