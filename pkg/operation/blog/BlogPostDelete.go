package blog

import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
)

func BlogPostDelete(id int64,email string,role string)  (*model.Blog,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	blogData := model.Blog{}
	blogData.PreloadBlog(db).First(&blogData,id)
	if blogData.User.ID == user.ID || strings.EqualFold(env.ADMIN,role){
		err := db.Delete(blogData).Error
		if err!=nil{
			return nil,err
		}else{
			return &blogData,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}