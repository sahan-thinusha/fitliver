package blog

import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BlogCommentUpdate(comment *model.BlogComment,email string)  (*model.BlogComment,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	comment.User = user
	blogComment := model.BlogComment{}
	blogComment.PreloadBlogComment(db).First(&blogComment,comment.ID)

	if blogComment.User.ID == user.ID{
		blogComment.Comment = comment.Comment
		err := db.Save(blogComment).Error
		if err!=nil{
			return nil,err
		}else{
			return &blogComment,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}
}