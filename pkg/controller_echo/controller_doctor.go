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

func GetDoctors(c echo.Context) error {
	result, err := doctor.GetDoctors(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetNewDoctors(c echo.Context) error {
	result, err := doctor.GetNewDoctors(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
func ApproveDoctor(c echo.Context) error {
	result, err := doctor.ApproveDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerDoctor(g *echo.Group) {
	g.GET("api/doctor/:id", GetDoctor)
	g.GET("api/doctors", GetDoctors)
	g.GET("api/doctors/new", GetNewDoctors)
	g.POST("api/doctors/approve/:id", ApproveDoctor)
}
