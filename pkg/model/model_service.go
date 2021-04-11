package model

import "github.com/jinzhu/gorm"

type Service struct {
	Model
	Name string   `json:"name"`
	ConsultationService   []*ConsultationService  `gorm:"many2many:consultation_service_service" json:"consultation_service"`
}

func (Service) TableName() string {
	return "service"
}
func (m *Service) PreloadService(db *gorm.DB) *gorm.DB {
	return db
}
