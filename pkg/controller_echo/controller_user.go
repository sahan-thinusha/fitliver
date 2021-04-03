package controller_echo

import (
	"fitliver/pkg/api_echo/doctor"
	"fitliver/pkg/api_echo/patient"
	"fitliver/pkg/api_echo/user"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func CreateDoctor(c echo.Context) error {
	result, err := doctor.CreateDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func CreatePatient(c echo.Context) error {
	result, err := patient.CreatePatient(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func LoginController(g *echo.Group) {
	g.POST("api/user/authenticate", user.Login)
	g.POST("api/user/doctor/register",CreateDoctor)
	g.POST("api/user/patient/register",CreatePatient)

}


