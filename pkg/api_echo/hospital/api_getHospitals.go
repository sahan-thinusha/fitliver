package hospital

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/hospital"
	"fmt"
	"github.com/labstack/echo/v4"
)

func GetHospitals(c echo.Context) ([]*model.Hospital, error) {
	fmt.Println(c.Request().Header)
	result,err := op.GetHospitals()
	return result, err

}