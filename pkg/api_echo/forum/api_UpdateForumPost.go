package forum


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/forum"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UpdateFormReply(c echo.Context) (*model.ForumReply, error) {
	comment := model.ForumReply{}
	if error := c.Bind(&comment); error != nil {
		return nil, error
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)


	result,err := op.ForumReplyUpdate(&comment,claims.Sub)
	return result, err
}
