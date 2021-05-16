package availability

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/availability"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetAvailabilityByDateDoctor(c echo.Context) ([]*model.Availability, error) {
	date := c.QueryParam("date")
	id := c.Param("id")
	doctorID,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.GetAvailabilityByDateDoctor(date,doctorID)
	return result, err
}