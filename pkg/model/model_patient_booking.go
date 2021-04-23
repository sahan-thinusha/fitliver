package model


import (
	"github.com/jinzhu/gorm"
	"time"
)

type Patient_Booking struct{
	Model
	Patient    *Patient   `gorm:"foreignkey:patientID" json:"patient"`
	PatientID  int64
	ConsultationService    *ConsultationService   `gorm:"foreignkey:consultations_serviceID" json:"consultations_service"`
	ConsultationServiceID  int64
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
func (m *Patient_Booking) PreloadPPatient_Booking(db *gorm.DB) *gorm.DB {
	return db.Preload("Patient").Preload("ConsultationService").Preload("Doctor")
}
