package model

import (
	"github.com/jinzhu/gorm"
)

type Doctor struct {
	Model
	Name           string      `json:"name"`
	ContactNo      []ContactNo `json:"contactno"`
	Address        string      `json:"address"`
	Hospital       []Hospital  `json:"hospital"`
	Specialization string      `json:"specialization"`
	Gender         string      `json:"gender"`
	DateOfBirth    string      `json:"dob"`
	ProfilePic string  `json:"profilepic"`
}

func (Doctor) TableName() string {
	return "doctor"
}
func (m *Doctor) PreloadDoctor(db *gorm.DB) *gorm.DB {
	return db
}
