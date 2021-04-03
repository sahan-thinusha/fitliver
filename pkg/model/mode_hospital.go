package model

import "github.com/jinzhu/gorm"

type Hospital struct {
	Model
	Name string `json:"name"`
	Address string `json:"address"`
	Doctor       []*Doctor  `gorm:"many2many:doctor_hospital" json:"doctor"`
}

func (Hospital) TableName() string {
	return "hospital"
}
func (m *Hospital) PreloadHospital(db *gorm.DB) *gorm.DB {
	return db
}
