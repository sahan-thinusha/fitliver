package forum

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/forum"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetForumPostbyId(c echo.Context) (*model.ForumPost, error) {
	id := c.Param("id")
	forumId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.ForumPostbyId(forumId)
	return result, err
}
