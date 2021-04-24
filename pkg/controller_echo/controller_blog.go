package controller_echo


import (
	"fitliver/pkg/api_echo/blog"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

func CreateBlogPost(c echo.Context) error {
	result, err := blog.CreateBlogPost(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateBlogPost(c echo.Context) error {
	result, err := blog.UpdateBlogPost(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func DeleteBlogPost(c echo.Context) error {
	result, err := blog.DeleteBlogPost(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetBlogPost(c echo.Context) error {
	result, err := blog.GetBlogPostbyId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetBlogPosts(c echo.Context) error {
	result, err := blog.GetBlogPosts(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func CreateBlogComment(c echo.Context) error {
	result, err := blog.CreateBlogComment(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateBlogComment(c echo.Context) error {
	result, err := blog.UpdateBlogComment(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
func DeleteBlogComment(c echo.Context) error {
	result, err := blog.DeleteBlogComment(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func APIControllerBlog(g *echo.Group) {
	g.POST("api/blog", CreateBlogPost)
	g.PUT("api/blog/:id", UpdateBlogPost)
	g.DELETE("api/blog/:id", DeleteBlogPost)
	g.GET("api/blog/:id", GetBlogPost)
	g.GET("api/blog", GetBlogPosts)
	g.POST("api/blog/:id/comment", CreateBlogComment)
	g.PUT("api/blog/comment", UpdateBlogComment)
	g.DELETE("api/blog/comment/:id", DeleteBlogComment)
}
