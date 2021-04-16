package controller_echo

import (
	"fitliver/pkg/api_echo/healthrecord"
	"github.com/labstack/echo/v4"
	"net/http"
)


func SaveHealthRecord(c echo.Context) error {
	result, err := healthrecord.SaveHealthRecord(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateHealthRecord(c echo.Context) error {
	result, err := healthrecord.UpdateHealthRecord(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetHealthRecordsByPatient(c echo.Context) error {
	result, err := healthrecord.GetHealthRecordByPatient(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerHealthRecordService(g *echo.Group) {
	g.POST("api/health_record", SaveHealthRecord)
	g.PUT("api/health_record", UpdateHealthRecord)
	g.GET("api/health_record/:id", GetHealthRecordsByPatient)
}
