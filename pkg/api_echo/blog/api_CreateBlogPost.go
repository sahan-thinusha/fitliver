package blog

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/blog"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"strconv"
	"time"
)

func CreateBlogPost(c echo.Context) (*model.Blog, error) {
	title := c.FormValue("title")
	body := c.FormValue("body")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)


	blog := model.Blog{}
	blog.Title = title
	blog.Body = body

	file, err := c.FormFile("image")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return nil,err
		}
		defer src.Close()


		folderPath := "blog_images"
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			_ = os.Mkdir(folderPath, os.ModePerm)
		}

		fileSuffix := strconv.FormatInt(time.Now().UnixNano(), 10)
		dst, err := os.Create(folderPath + "/" + fileSuffix + "_" + file.Filename)
		if err != nil {
			return nil,err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return nil,err
		}

		blog.Image = dst.Name()
	}


	result,err := op.BlogPostCreate(&blog,claims.Sub)
	return result, err
}
