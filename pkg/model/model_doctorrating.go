package model



import (
	"github.com/jinzhu/gorm"
)

type Rating struct {
	Model
	Doctor     *Doctor `gorm:"foreignkey:doctorID" json:"doctor"`
	DoctorID int64
	Rate float64  `json:"rating"`
	User     *User `gorm:"foreignkey:userID" json:"user"`
	UserID int64
}

func (Rating) TableName() string {
	return "rating"
}
func (m *Rating) PreloadRating(db *gorm.DB) *gorm.DB {
	return  db.Preload("User").Preload("User.Doctor")
}
