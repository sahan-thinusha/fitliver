package controller_echo


import (
	"fitliver/pkg/api_echo/workout"
	"github.com/labstack/echo/v4"
	"net/http"
)


func SaveWorkoutPlan(c echo.Context) error {
	result, err := workout.SaveWorkoutPlan(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateWorkoutPlan(c echo.Context) error {
	result, err := workout.UpdateWorkoutPlan(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetWorkoutPlansByPatient(c echo.Context) error {
	result, err := workout.GetPatientWorkoutPlans(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetWorkouts(c echo.Context) error {
	result, err := workout.GetWorkouts(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func APIControllerWorkoutPlan(g *echo.Group) {
	g.POST("api/workout", SaveWorkoutPlan)
	g.PUT("api/workout", UpdateWorkoutPlan)
	g.GET("api/workout/:id", GetWorkoutPlansByPatient)
	g.GET("api/workouts", GetWorkouts)

}
