package controller_echo



import (
	"fitliver/pkg/api_echo/availability"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func CreateAvailability(c echo.Context) error {
	result, err := availability.AddAvailability(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateAvailability(c echo.Context) error {
	result, err := availability.UpdateAvailability(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func DeleteAvailability(c echo.Context) error {
	result, err := availability.DeleteAvailability(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetAvailability(c echo.Context) error {
	result, err := availability.GetAvailability(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetAvailabilitys(c echo.Context) error {
	result, err := availability.GetAvailabilityByDate(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func APIControllerAvailability(g *echo.Group) {
	g.POST("api/availability", CreateAvailability)
	g.PUT("api/availability/:id", UpdateAvailability)
	g.DELETE("api/availability/:id", DeleteAvailability)
	g.GET("api/availability/:id", GetAvailability)
	g.GET("api/availability", GetAvailabilitys)

}
