package model

import "github.com/jinzhu/gorm"

type ContactNo struct {
	Model
	ContactNo string `json:"contactno"`
	Doctor     Doctor `gorm:"foreignkey:doctorID" json:"doctor"`
	DoctorID int64
}

func (ContactNo) TableName() string {
	return "contactno"
}
func (m *ContactNo) PreloadContactNo(db *gorm.DB) *gorm.DB {
	return db
}
