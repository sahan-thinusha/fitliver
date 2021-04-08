package blog

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BlogCommentCreate(comment *model.BlogComment,email string)  (*model.BlogComment,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	comment.User = user
	blog := model.Blog{}
	db.Model(model.Blog{}).First(&blog,comment.BlogID)
	comment.Blog = blog
	err := db.Create(comment).Error
	if err!=nil{
		return nil,err
	}else{
		return comment,nil
	}
}