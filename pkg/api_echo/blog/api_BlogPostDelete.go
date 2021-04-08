package blog

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/blog"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func DeleteBlogPost(c echo.Context) (*model.Blog, error) {
	id := c.Param("id")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	blogId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.BlogPostDelete(blogId,claims.Sub,claims.Auth)
	return result, err
}
