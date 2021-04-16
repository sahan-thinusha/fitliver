package forum



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ForumPostCreate(forum *model.ForumPost,email string)  (*model.ForumPost,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	forum.User = &user
	err := db.Create(forum).Error
	if err!=nil{
		return nil,err
	}else{
		return forum,nil
	}
}