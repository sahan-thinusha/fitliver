package hospital


import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/hospital"
	"github.com/labstack/echo/v4"
)

func SaveHospital(c echo.Context) (*model.Hospital, error) {
	hospital := model.Hospital{}
	if err := c.Bind(&hospital); err != nil {
		return nil, err
	}

	result,err := op.AddHospital(&hospital)
	return result, err
}
