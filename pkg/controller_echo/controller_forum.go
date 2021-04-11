package controller_echo

import (
	"fitliver/pkg/api_echo/forum"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateForumPost(c echo.Context) error {
	result, err := forum.CreateForumPost(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateForumPost(c echo.Context) error {
	result, err := forum.UpdateForumPost(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func DeleteForumPost(c echo.Context) error {
	result, err := forum.DeleteForumPost(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetForumPost(c echo.Context) error {
	result, err := forum.GetForumPostbyId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetForumPosts(c echo.Context) error {
	result, err := forum.GetForumPosts(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func CreateForumReply(c echo.Context) error {
	result, err := forum.CreateForumReply(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func UpdateForumReply(c echo.Context) error {
	result, err := forum.UpdateFormReply(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
func DeleteForumReply(c echo.Context) error {
	result, err := forum.CreateForumPost(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func APIControllerForum(g *echo.Group) {
	g.POST("api/forum", CreateForumPost)
	g.PUT("api/forum/:id", UpdateForumPost)
	g.DELETE("api/forum/:id", DeleteForumPost)
	g.GET("api/forum/:id", GetForumPost)
	g.GET("api/forum", GetForumPosts)
	g.POST("api/forum/:id/reply", CreateForumReply)
	g.PUT("api/forum/reply", UpdateForumReply)
	g.DELETE("api/forum/reply/:id", DeleteForumReply)
}
