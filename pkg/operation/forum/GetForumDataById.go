package forum


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ForumPostbyId(id int64)  (*model.ForumPost,error){
	db :=env.RDB
	forum := model.ForumPost{}
	forum.PreloadForumPost(db).First(&forum,id)
	return &forum,nil
}