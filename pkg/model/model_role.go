package model
import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	Model
	Name string `json:"name"`

	User []*User `gorm:"many2many:user_role" json:"user"`
}

func (m *Role) PreloadRole(db *gorm.DB) *gorm.DB {
	return db.Preload("User")
}