package blog

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/blog"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UpdateBlogComment(c echo.Context) (*model.BlogComment, error) {
	comment := model.BlogComment{}
	if error := c.Bind(&comment); error != nil {
		return nil, error
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)


	result,err := op.BlogCommentUpdate(&comment,claims.Sub)
	return result, err
}
