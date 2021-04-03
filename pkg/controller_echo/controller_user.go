package controller_echo

import (
	"fitliver/pkg/api_echo/doctor"
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

func LoginController(g *echo.Group) {
	g.POST("/user/authenticate", user.Login)
	g.POST("/user/doctor/register",CreateDoctor)
}


