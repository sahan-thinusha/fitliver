package model

import "github.com/jinzhu/gorm"

type Availability struct {
	Model
	Doctor     *Doctor `gorm:"foreignkey:doctorID" json:"doctor"`
	DoctorID int64
	Hospital     *Hospital `gorm:"foreignkey:hospitalID" json:"hospital"`
	HospitalID int64
	AvailableDate string `json:"available_date"`
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
	IsFree bool `json:"isfree"`

}

func (Availability) TableName() string {
	return "availability"
}
func (m *Availability) PreloadAvailability(db *gorm.DB) *gorm.DB {
	return db.Preload("Doctor").Preload("Doctor.User")
}