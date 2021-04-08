package blog

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BlogPostbyId(id int64)  (*model.Blog,error){
	db :=env.RDB
	blogData := model.Blog{}
	blogData.PreloadBlog(db).First(&blogData,id)
	return &blogData,nil
}