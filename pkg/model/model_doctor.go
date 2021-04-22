package model

import (
	"github.com/jinzhu/gorm"
)

type Doctor struct {
	Model
	Name           string      `json:"name"`
	ContactNo      []*ContactNo `json:"contactno"`
	Address        string      `json:"address"`
	Hospital       []*Hospital  `gorm:"many2many:doctor_hospital" json:"hospital"`
	Specialization string      `json:"specialization"`
	Gender         string      `json:"gender"`
	DateOfBirth    string      `json:"dob"`
	ProfilePic string  `json:"profilepic"`
	User     User `gorm:"foreignkey:userID" json:"user"`
	Status  string
	UserID int64
	Rate float64  `json:"rating"`
}

func (Doctor) TableName() string {
	return "doctor"
}
func (m *Doctor) PreloadDoctor(db *gorm.DB) *gorm.DB {
	return db.Preload("Hospital").Preload("ContactNo")
}
