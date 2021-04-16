package model

import (
	"github.com/jinzhu/gorm"
)

type Patient struct {
	Model
	Name           string      `json:"name"`
	Gender         string      `json:"gender"`
	DateOfBirth    string      `json:"dob"`
	ProfilePic string  `json:"profilepic"`
	User     *User `gorm:"foreignkey:userID" json:"user"`
	UserID int64
}

func (Patient) TableName() string {
	return "patient"
}
func (m *Patient) PreloadPatient(db *gorm.DB) *gorm.DB {
	return db
}
