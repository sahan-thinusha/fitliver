package controller_echo


import (
	"fitliver/pkg/api_echo/doctor"
	"fitliver/pkg/api_echo/hospital"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func GetHospitals(c echo.Context) error {
	result, err := hospital.GetHospitals(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func GetHospitalById(c echo.Context) error {
	result, err := doctor.GetDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
func APIControllerHospital(g *echo.Group) {
	g.GET("/hospital/:id", GetHospitalById)
	g.GET("/hospital", GetHospitals)
}

func APIControllerHospitalBase(g *echo.Group) {
	g.GET("/hospital", GetHospitals)
}
