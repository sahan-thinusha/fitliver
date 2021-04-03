package controller_echo

import (
	"fitliver/pkg/api_echo/doctor"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func GetDoctor(c echo.Context) error {
	result, err := doctor.GetDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerDoctor(g *echo.Group) {
	g.GET("/doctor/:id", GetDoctor)
}
