package model


import (
	"github.com/jinzhu/gorm"
	"time"
)

type Patient_Booking struct{
	Model
	Patient    *Patient   `gorm:"foreignkey:patientID" json:"patient"`
	PatientID  int64
	BookingService    *BookingService   `gorm:"foreignkey:booking_serviceID" json:"booking_service"`
	BookingServiceID  int64
	PurchasedAt time.Time `json:"purchased_at"`
	BookedDate string
	TimeFrom string
	TimeTo string
	Hospital     *Hospital `gorm:"foreignkey:hospitalID" json:"hospital"`
	HospitalID int64
}

func (Patient_Booking) TableName() string {
	return "patient_booking"
}
func (m *Patient_Booking) PreloadPatient_Booking(db *gorm.DB) *gorm.DB {
	return db.Preload("Patient").Preload("BookingService").Preload("BookingService.Doctor")
}
