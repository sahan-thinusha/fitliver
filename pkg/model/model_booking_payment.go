package model


import (
	"github.com/jinzhu/gorm"
)

type BookingPayment struct {
	Model
	Amount   int64      `json:"amount"`
	ReceiptEmail string `json:"receiptMail"`
	Name string   `json:"name"`
	PaymentId string   `json:"payment_id"`
	Patient     Patient `gorm:"foreignkey:patientID" json:"patient"`
	Status  string   `json:"status"`
	PatientID int64
	Booking     BookingService `gorm:"foreignkey:bookingID" json:"booking"`
	BookingID int64
	BookingRequest     BookingRequest `gorm:"foreignkey:bookingRequestID" json:"booking_request"`
	BookingRequestID int64
}

func (BookingPayment) TableName() string {
	return "booking_payment"
}
func (m *BookingPayment) PreloadBookingPayment(db *gorm.DB) *gorm.DB {
	return db
}