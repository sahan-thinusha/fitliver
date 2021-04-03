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
	g.POST("api/blog", CreateDoctor)
	g.GET("api/blog/:id", GetDoctor)
	g.GET("api/blog", GetDoctor)
	g.POST("api/blog/comment:id", GetDoctor)
	g.PUT("api/blog/comment", GetDoctor)
	g.DELETE("api/blog/comment", GetDoctor)
}
