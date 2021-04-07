package blog

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BlogPostCreate(blog *model.Blog,email string)  (*model.Blog,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&model.User{})
	blog.User = user
	err := db.Create(blog).Error
	if err!=nil{
		return nil,err
	}else{
		return blog,nil
	}
}