package model

import "github.com/jinzhu/gorm"

type ContactNo struct {
	Model
	ContactNo string `json:"contactno"`
}

func (ContactNo) TableName() string {
	return "contactno"
}
func (m *ContactNo) PreloadContactNo(db *gorm.DB) *gorm.DB {
	return db
}
