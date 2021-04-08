package model

import "github.com/jinzhu/gorm"

type BlogComment struct {
	Model
	Comment string `json:"comment"  gorm:"size:2000"`
	User    User   `gorm:"foreignkey:userID" json:"user"`
	UserID  int64
	Blog    Blog   `gorm:"foreignkey:blogID" json:"blog"`
	BlogID  int64
}

func (BlogComment) TableName() string {
	return "blog_comment"
}
func (m *BlogComment) PreloadBlogComment(db *gorm.DB) *gorm.DB {
	return db.Preload("User").Preload("Blog")
}
