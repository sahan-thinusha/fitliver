package forum


import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
)

func ForumPostDelete(id int64,email string,role string)  (*model.ForumPost,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	forumData := model.ForumPost{}
	forumData.PreloadForumPost(db).First(&forumData,id)
	if forumData.User.ID == user.ID || strings.EqualFold(env.ADMIN,role){
		err := db.Delete(forumData).Error
		if err!=nil{
			return nil,err
		}else{
			return &forumData,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}