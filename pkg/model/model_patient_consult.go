package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Patient_Consult struct{
	Model
	Patient    *Patient   `gorm:"foreignkey:patientID" json:"patient"`
	PatientID  int64
	ConsultationService    *ConsultationService   `gorm:"foreignkey:consultations_serviceID" json:"consultations_service"`
	ConsultationServiceID  int64
	PurchasedAt time.Time `json:"purchased_at"`
	Duration int64
}

func (Patient_Consult) TableName() string {
	return "patient_consult"
}
func (m *Patient_Consult) PreloadPatient_Consult(db *gorm.DB) *gorm.DB {
	return db.Preload("Patient").Preload("ConsultationService").Preload("Doctor")
}
