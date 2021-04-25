package model

import (
	"github.com/jinzhu/gorm"
)

type ConsultationRequest struct{
	Model
	Patient    *Patient   `gorm:"foreignkey:patientID" json:"patient"`
	PatientID  int64
	Package    *Package   `gorm:"foreignkey:packageID" json:"package"`
	PackageID  int64
	Status string
}

func (ConsultationRequest) TableName() string {
	return "consultation_request"
}
func (m *ConsultationRequest) PreloadConsultationRequest(db *gorm.DB) *gorm.DB {
	return db.Preload("Patient").Preload("Package")
}
