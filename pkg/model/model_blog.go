package model

import "github.com/jinzhu/gorm"

type Blog struct {
	Model
	Title string `json:"title"`
	Body  string `json:"body"  gorm:"size:2000"`
	Image string `json:"image" gorm:"size:256"`
	User     User `gorm:"foreignkey:userID" json:"user"`
	UserID int64
	BlogComment       []*BlogComment  `json:"blog_comment"`


}

func (Blog) TableName() string {
	return "blog"
}
func (m *Blog) PreloadBlog(db *gorm.DB) *gorm.DB {
	return db.Preload("User").Preload("BlogComment")
}
