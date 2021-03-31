package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Token     string  `json:"token"`
	Role      []*Role `gorm:"many2many:user_role" json:"role"`
}

func (m *User) PreloadUser(db *gorm.DB) *gorm.DB {
	return db.Preload("Role")
}