package model

import "time"

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Tabler interface {
	TableName() string
}

func InitModels(db *gorm.DB) {
	db.AutoMigrate(&Doctor{})
	db.AutoMigrate(&Patient{})
	db.AutoMigrate(&ContactNo{})
	db.AutoMigrate(&Hospital{})
	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&BlogComment{})
	db.AutoMigrate(&ForumPost{})
	db.AutoMigrate(&ForumReply{})
	db.AutoMigrate(&User{})
}
