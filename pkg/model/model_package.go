package model

import "github.com/jinzhu/gorm"

type Package struct {
	Model
	Name string `json:"name"`
	Amount  int64 `json:"amount"`
	Description  string `json:"description"`
	ConsultationService     *ConsultationService `gorm:"foreignkey:consultationserviceID" json:"consultationservice"`
	ConsultationServiceID int64
	Duration float64 `json:"duration"`
}

func (Package) TableName() string {
	return "package"
}
func (m *Package) PreloadPackage(db *gorm.DB) *gorm.DB {
	return db.Preload("ConsultationService")
}
