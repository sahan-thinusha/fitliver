package controller_echo

import (
	"fitliver/pkg/api_echo/consultation_service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateConsultationService(c echo.Context) error {
	result, err := consultation_service.CreateConsultationService(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateConsultationService(c echo.Context) error {
	result, err := consultation_service.UpdateConsultationService(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func DeleteConsultationService(c echo.Context) error {
	result, err := consultation_service.DeleteConsultationService(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}




func GetAllConsultationServices(c echo.Context) error {
	result, err := consultation_service.GetAllConsultationServices(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func GetConsultationServicesById(c echo.Context) error {
	result, err := consultation_service.GetConsultationServicesById(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func GetAllConsultationServicesForDoctor(c echo.Context) error {
	result, err := consultation_service.GetAllConsultationServicesForDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}




func APIControllerConsultationService(g *echo.Group) {
	g.POST("api/consultationservice", CreateConsultationService)
	g.PUT("api/consultationservice", UpdateConsultationService)
	g.DELETE("api/consultationservice/:id", DeleteConsultationService)
	g.GET("api/consultationservice/:id", GetConsultationServicesById)
	g.GET("api/consultationservice", GetAllConsultationServices)
	g.GET("api/consultationservice/doctor", GetAllConsultationServicesForDoctor)
}
