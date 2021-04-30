package controller_echo

import (
	"fitliver/pkg/api_echo/patient"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetPatient(c echo.Context) error {
	result, err := patient.GetPatient(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetPatients(c echo.Context) error {
	result, err := patient.GetPatients(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetPatientsbyDoctor(c echo.Context) error {
	result, err := patient.GetPDoctorSubscribedPatients(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func APIControllerPatient(g *echo.Group) {
	g.GET("api/patient/:id", GetPatient)
	g.GET("api/patients", GetPatients)
	g.GET("api/doctor/patients", GetPatientsbyDoctor)

}
