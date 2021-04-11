package forum

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/forum"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateForumPost(c echo.Context) (*model.ForumPost, error) {
	title := c.FormValue("title")
	body := c.FormValue("body")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)


	forum := model.ForumPost{}
	forum.Title = title
	forum.Body = body



	result,err := op.ForumPostCreate(&forum,claims.Sub)
	return result, err
}
