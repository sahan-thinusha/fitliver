package fooddata

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/diet_plan"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetPatientDietPlan(c echo.Context) ([]*model.DietPlan, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	result,err := op.GetPatientDietPlans(id)
	return result, err
}