package blog

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/blog"
	"github.com/labstack/echo/v4"
)

func GetBlogPosts(c echo.Context) ([]*model.Blog, error) {
	result,err := op.BlogPosts()
	return result, err
}
