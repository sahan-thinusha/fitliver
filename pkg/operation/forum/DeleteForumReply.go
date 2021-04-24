package forum

import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
)

func ForumReplyDelete(id int64,email string,role string)  (*model.ForumReply,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	forumReply := model.ForumReply{}
	forumReply.PreloadForumReply(db).First(&forumReply,id)
	if forumReply.User.ID == user.ID || strings.EqualFold(env.ADMIN,role){
		err := db.Delete(&forumReply).Error
		if err!=nil{
			return nil,err
		}else{
			return &forumReply,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}