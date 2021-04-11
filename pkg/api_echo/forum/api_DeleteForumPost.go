package forum

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/forum"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func DeleteForumPost(c echo.Context) (*model.ForumPost, error) {
	id := c.Param("id")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	forumId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.ForumPostDelete(forumId,claims.Sub,claims.Auth)
	return result, err
}
