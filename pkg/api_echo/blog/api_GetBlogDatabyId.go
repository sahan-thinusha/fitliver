package blog

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/blog"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetBlogPostbyId(c echo.Context) (*model.Blog, error) {
	id := c.Param("id")
		blogId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.BlogPostbyId(blogId)
	return result, err
}
