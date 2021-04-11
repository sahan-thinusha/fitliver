package forum

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/forum"
	"github.com/labstack/echo/v4"
)

func GetForumPosts(c echo.Context) ([]*model.ForumPost, error) {
	result,err := op.ForumPosts()
	return result, err
}
