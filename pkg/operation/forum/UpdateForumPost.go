package forum

import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ForumPostUpdate(forum *model.ForumPost,email string)  (*model.ForumPost,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	forumData := model.ForumPost{}
	forum.PreloadForumPost(db).First(&forumData,forum.ID)
	if forumData.User.ID == user.ID{
		forumData.Body = forum.Body
		forumData.Title = forum.Title
		err := db.Save(forumData).Error
		if err!=nil{
			return nil,err
		}else{
			return &forumData,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}