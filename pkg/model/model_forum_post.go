package model

import "github.com/jinzhu/gorm"

type ForumPost struct {
	Model
	Title string `json:"title"`
	Body  string `json:"body"  gorm:"size:2000"`
	User     *User `gorm:"foreignkey:userID" json:"user"`
	HashTags string `json:"tags"`
	UserID int64
	ForumReply       []*ForumReply  `json:"forum_reply"`
}



func (ForumPost) TableName() string {
	return "forum_post"
}
func (m *ForumPost) PreloadForumPost(db *gorm.DB) *gorm.DB {
	return db.Preload("User").Preload("ForumReply").Preload("ForumReply.User")
}
