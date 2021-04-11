package forum


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ForumPosts() ([]*model.ForumPost,error){
	db :=env.RDB
	forumData := model.ForumPost{}
	forums := []*model.ForumPost{}
	forumData.PreloadForumPost(db).Find(&forums)
	return forums,nil
}