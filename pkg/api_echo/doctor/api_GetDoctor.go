package doctor

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/doctor"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetDoctor(c echo.Context) (*model.Doctor, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	result,err := op.GetDoctor(id)
	return result, err

}