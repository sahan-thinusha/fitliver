package model

import "github.com/jinzhu/gorm"

type HealthRecord struct {
	Model
	Progress string `json:"progress"`
	ASTProgress string `json:"astProgress"`
	AST float64 `json:"ast"`
	ALTProgress string `json:"altProgress"`
	ALT float64 `json:"alt"`
	FATProgress string `json:"fatProgress"`
	FAT float64 `json:"fat"`
	Patient    *Patient   `gorm:"foreignkey:patientID" json:"patient"`
	PatientID  int64
	Doctor   *Doctor   `gorm:"foreignkey:doctorID" json:"doctor"`
	DoctorID  int64
}
func (HealthRecord) TableName() string {
	return "health_record"
}

func (m *HealthRecord) PreloadHealthRecord(db *gorm.DB) *gorm.DB {
	return db.Preload("Doctor").Preload("Patient")
}