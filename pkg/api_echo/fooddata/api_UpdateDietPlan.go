package fooddata

import (
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/diet_plan"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

func UpdateDietPlan(c echo.Context) (*model.DietPlan, error) {
	dietPlan := model.DietPlan{}
	date := c.QueryParam("date")
	if error := c.Bind(&dietPlan); error != nil {
		return nil, error
	}

	planDate, _ := time.Parse(Time, date)

	dietPlan.PlanDate = &planDate
	result,err := op.DietPlanUpdate(&dietPlan)
	return result, err
}
