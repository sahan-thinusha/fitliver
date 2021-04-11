package forum

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func FormReplyCreate(reply *model.ForumReply,email string)  (*model.ForumReply,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	reply.User = user
	forum := model.ForumPost{}
	db.Model(model.ForumPost{}).First(&forum,reply.ForumPostID)
	reply.ForumPost = forum
	err := db.Create(reply).Error
	if err!=nil{
		return nil,err
	}else{
		return reply,nil
	}
}