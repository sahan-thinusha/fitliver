package doctor

import (
	model "fitliver/pkg/model"
	op "fitliver/pkg/operation/doctor"
)

import (
	"github.com/labstack/echo/v4"
)

func CreateDoctor(c echo.Context) (*model.Doctor, error) {
	doctor := model.Doctor{}
	if error := c.Bind(&doctor); error != nil {
		return nil, error
	}
	result,err := op.DoctorRegister(&doctor)

	return result, err

}
