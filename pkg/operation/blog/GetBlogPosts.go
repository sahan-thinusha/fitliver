package blog


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BlogPosts()  ([]*model.Blog,error){
	db :=env.RDB
	blogData := model.Blog{}
	blogs := []*model.Blog{}
	blogData.PreloadBlog(db).Find(&blogs)
	return blogs,nil
}