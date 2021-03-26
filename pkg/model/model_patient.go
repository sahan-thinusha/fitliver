package model

import (
	"github.com/jinzhu/gorm"
)

type Patient struct {
	Model
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

func (Patient) TableName() string {
	return "patient"
}
func (m *Patient) PreloadPatient(db *gorm.DB) *gorm.DB {
	return db
}
