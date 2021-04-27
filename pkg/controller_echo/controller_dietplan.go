package controller_echo

import (
	"fitliver/pkg/api_echo/fooddata"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetDietPlan(c echo.Context) error {
	result, err := fooddata.GetFoodData(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func SaveDietPlan(c echo.Context) error {
	result, err := fooddata.SaveDietPlan(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func UpdateDietPlan(c echo.Context) error {
	result, err := fooddata.UpdateDietPlan(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func GetSavedPatientDietPlans(c echo.Context) error {
	result, err := fooddata.GetPatientDietPlan(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func GetSavedPatientDietPlansForDates(c echo.Context) error {
	result, err := fooddata.GetPatientDietPlanForDateRange(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerDietPlanService(g *echo.Group) {
	g.POST("api/diet_plan", GetDietPlan)
	g.POST("api/diet_plan", SaveDietPlan)
	g.PUT("api/diet_plan", UpdateDietPlan)
	g.GET("api/diet_plan/:id", GetSavedPatientDietPlans)
	g.GET("api/diet_plan/:id/by_date", GetSavedPatientDietPlansForDates)
}
