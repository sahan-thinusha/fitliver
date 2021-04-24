package model

import "github.com/jinzhu/gorm"

type ForumReply struct {
	Model
	Comment string `json:"comment"  gorm:"size:2000"`
	User    User   `gorm:"foreignkey:userID" json:"user"`
	UserID  int64
	ForumPost    ForumPost   `gorm:"foreignkey:forumpostID" json:"forumPost"`
	ForumPostID  int64
}

func (ForumReply) TableName() string {
	return "forum_reply"
}
func (m *ForumReply) PreloadForumReply(db *gorm.DB) *gorm.DB {
	return db.Preload("User").Preload("ForumPost")
}