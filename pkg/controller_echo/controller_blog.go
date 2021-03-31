package controller_echo


import (
	"fitliver/pkg/api_echo/doctor"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func CreateBlogPost(c echo.Context) error {
	result, err := doctor.CreateDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func APIControllerBlog(g *echo.Group) {
	g.POST("/doctor/create", CreateDoctor)
	g.GET("/doctor/:id", GetDoctor)
}
