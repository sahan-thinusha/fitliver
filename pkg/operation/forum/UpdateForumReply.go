package forum



import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ForumReplyUpdate(forumReply *model.ForumReply,email string)  (*model.ForumReply,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	forumReply.User = user
	ForumReplyData := model.ForumReply{}
	ForumReplyData.PreloadForumReply(db).First(&ForumReplyData,forumReply.ID)
	if ForumReplyData.User.ID == user.ID{
		ForumReplyData.Comment = forumReply.Comment
		err := db.Save(ForumReplyData).Error
		if err!=nil{
			return nil,err
		}else{
			return &ForumReplyData,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}
}