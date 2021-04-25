package model

import (
	"github.com/jinzhu/gorm"
)

type BookingRequest struct{
	Model
	Patient    *Patient   `gorm:"foreignkey:patientID" json:"patient"`
	PatientID  int64
	Booking     *BookingService `gorm:"foreignkey:bookingID" json:"booking"`
	BookingID int64
	Status string
	BookedDate string
	TimeFrom string
	TimeTo string
	Hospital     *Hospital `gorm:"foreignkey:hospitalID" json:"hospital"`
	HospitalID int64
}

func (BookingRequest) TableName() string {
	return "booking_request"
}
func (m *BookingRequest) PreloadBookingRequest(db *gorm.DB) *gorm.DB {
	return db.Preload("Patient").Preload("Booking").Preload("Hospital")
}
