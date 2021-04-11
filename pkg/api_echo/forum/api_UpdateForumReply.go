package forum



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/forum"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func UpdateForumPost(c echo.Context) (*model.ForumPost, error) {
	id := c.Param("id")
	title := c.FormValue("title")
	body := c.FormValue("body")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)


	forum := model.ForumPost{}
	forum.ID, _ = strconv.ParseInt(id, 10, 64)
	forum.Title = title
	forum.Body = body



	result,err := op.ForumPostUpdate(&forum,claims.Sub)
	return result, err
}
