package availability

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/availability"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetAvailability(c echo.Context) (*model.Availability, error) {
	id := c.Param("id")
	availabilityID,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.GetAvailability(availabilityID)
	return result, err
}
