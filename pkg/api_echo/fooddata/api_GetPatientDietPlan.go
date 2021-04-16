package fooddata

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/diet_plan"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetPatientDietPlan(c echo.Context) ([]*model.DietPlan, error) {
	id:= c.Param("id")
	patientId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.GetPatientDietPlans(patientId)
	return result, err
}