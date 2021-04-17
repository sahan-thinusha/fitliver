package controller_echo

import (
	"fitliver/pkg/api_echo/doctor"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddRating(c echo.Context) error {
	result, err := doctor.AddRating(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func UpdateRating(c echo.Context) error {
	result, err := doctor.UpdateRating(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func GetUserRatingForDoctor(c echo.Context) error {
	result, err := doctor.GetUserRatingForDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetOverallRatingForDoctor(c echo.Context) error {
	result, err := doctor.GetOverallRatingForDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}




func APIControllerRating(g *echo.Group) {
	g.POST("api/doctor/rating", AddRating)
	g.PUT("api/doctor/rating", UpdateRating)
	g.GET("api/doctor/rating/user/:id", GetUserRatingForDoctor)
	g.GET("api/doctor/rating/:id", GetOverallRatingForDoctor)
}
