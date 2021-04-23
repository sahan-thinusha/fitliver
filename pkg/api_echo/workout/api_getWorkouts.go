package workout

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/workout"
	"github.com/labstack/echo/v4"
)

func GetWorkouts(c echo.Context) ([]*model.Workout, error) {
	result,err := op.GetWorkouts()
	return result, err
}
