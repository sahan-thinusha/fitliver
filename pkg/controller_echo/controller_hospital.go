package controller_echo

import (
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

func SaveHospital(c echo.Context) error {
	result, err := hospital.SaveHospital(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
func APIControllerHospital(g *echo.Group) {
	g.GET("api/gethospitals", GetHospitals)
	g.POST("api/hospital", SaveHospital)

}

func APIControllerHospitalBase(g *echo.Group) {
	g.GET("api/hospital", GetHospitals)
}
