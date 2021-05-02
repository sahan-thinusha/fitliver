package model

import (
	"github.com/jinzhu/gorm"
)

type Payment struct {
	Model
	Amount   int64      `json:"amount"`
	ReceiptEmail string `json:"receiptMail"`
	Name string   `json:"name"`
	PaymentId string   `json:"payment_id"`
	Patient     *Patient `gorm:"foreignkey:patientID" json:"patient"`
	Status  string   `json:"status"`
	PatientID int64
	Package     Package `gorm:"foreignkey:packageID" json:"package"`
	PackageID int64
}

func (Payment) TableName() string {
	return "payment"
}
func (m *Payment) PreloadPayment(db *gorm.DB) *gorm.DB {
	return db
}
