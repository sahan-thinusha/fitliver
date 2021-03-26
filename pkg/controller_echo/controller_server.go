package controller_echo

import (
	apiecho "fitliver/pkg/api_echo"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func CreateDoctor(c echo.Context) error {
	result, err := apiecho.CreateDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetDoctor(c echo.Context) error {
	result, err := apiecho.GetDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerEcho(g *echo.Group) {
	g.POST("/api/doctor/create", CreateDoctor)
	g.GET("/api/doctor/:id", GetDoctor)

}
