package fooddata

import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/diet_plan"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

func SaveDietPlan(c echo.Context) (*model.DietPlan, error) {
	dietPlan := model.DietPlan{}
	date := c.QueryParam("date")
	if error := c.Bind(&dietPlan); error != nil {
		return nil, error
	}

	planDate, _ := time.Parse(Time, date)

	dietPlan.PlanDate = &planDate

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.DietPlanSave(&dietPlan,claims.Sub)
	return result, err
}
