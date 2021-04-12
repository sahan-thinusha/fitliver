package model

import "github.com/jinzhu/gorm"

type ConsultationService struct {
	Model
	Name     string     `json:"name"`
	Packages []*Package `json:"package"`
	Service  []*Service `gorm:"many2many:consultation_service_service" json:"service"`
	Doctor     *Doctor `gorm:"foreignkey:doctorID" json:"doctor"`
	DoctorID int64
}

func (ConsultationService) TableName() string {
	return "consultation_service"
}
func (m *ConsultationService) PreloadConsultationService(db *gorm.DB) *gorm.DB {
	return db.Preload("Doctor").Preload("User").Preload("Package").Preload("Service")
}

func (m *ConsultationService) PreloadConsultationServiceForUser(db *gorm.DB) *gorm.DB {
	return db.Preload("Doctor").Preload("User")
}
