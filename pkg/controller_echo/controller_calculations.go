package controller_echo


import (
	"fitliver/pkg/api_echo/other"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func CalculateBMI(c echo.Context) error {
	result, err := other.CalculateBMI(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerCalculation(g *echo.Group) {
	g.GET("/other/bmi", CalculateBMI)
}