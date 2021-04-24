package blog


import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BlogPostUpdate(blog *model.Blog,email string)  (*model.Blog,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	blogData := model.Blog{}
	blogData.PreloadBlog(db).First(&blogData,blog.ID)
	if blogData.User.ID == user.ID{
		blogData.Image = blog.Image
		blogData.Body = blog.Body
		blogData.Title = blog.Title
		err := db.Save(&blogData).Error
		if err!=nil{
			return nil,err
		}else{
			return &blogData,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}