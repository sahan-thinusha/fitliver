package blog

import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
)

func BlogCommentDelete(id int64,email string,role string)  (*model.BlogComment,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	blogComment := model.BlogComment{}
	blogComment.PreloadBlogComment(db).First(&blogComment,id)
	if blogComment.User.ID == user.ID || strings.EqualFold(env.ADMIN,role){
		err := db.Delete(blogComment).Error
		if err!=nil{
			return nil,err
		}else{
			return &blogComment,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}