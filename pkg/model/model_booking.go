package model

import "github.com/jinzhu/gorm"

type BookingService struct {
	Model
	Name     string     `json:"name"`
	Amount  int64 `json:"amount"`
	Doctor     *Doctor `gorm:"foreignkey:doctorID" json:"doctor"`
	DoctorID int64
}

func (BookingService) TableName() string {
	return "booking_service"
}
func (m *BookingService) PreloadBookingService(db *gorm.DB) *gorm.DB {
	return db.Preload("Doctor")
}