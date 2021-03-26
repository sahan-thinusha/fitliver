package model

import "github.com/jinzhu/gorm"

type Hospital struct {
	Model
	Name string `json:"name"`
}

func (Hospital) TableName() string {
	return "hospital"
}
func (m *Hospital) PreloadHospital(db *gorm.DB) *gorm.DB {
	return db
}
