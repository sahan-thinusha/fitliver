package workout

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/workout"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UpdateWorkoutPlan(c echo.Context) (*model.WorkoutPlan, error) {
	workout := model.WorkoutPlan{}
	if err := c.Bind(&workout); err != nil {
		return nil, err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.UpdateWorkoutPlan(&workout,claims.Sub)
	return result, err
}
