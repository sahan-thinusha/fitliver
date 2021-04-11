package model

import "github.com/jinzhu/gorm"

type ConsultationService struct {
	Model
	Name     string     `json:"name"`
	Packages []*Package `json:"package"`
	Service  []*Service `gorm:"many2many:consultation_service_service" json:"service"`
}

func (ConsultationService) TableName() string {
	return "consultation_service"
}
func (m *ConsultationService) PreloadConsultationService(db *gorm.DB) *gorm.DB {
	return db
}
