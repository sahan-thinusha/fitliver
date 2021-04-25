package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type  DietPlan struct{
	Model
	Food      []*FoodData `json:"food" gorm:"foreignKey:DietPlanID"`
	Type  string `json:"type"`
	PlanDate *time.Time `json:"date"`
	MealName string `json:"mealName"`
	Patient     *Patient `gorm:"foreignkey:patientID" json:"patient"`
	PatientID int64
}

func (DietPlan) TableName() string {
	return "diet_plan"
}
func (m *DietPlan) PreloadDietPlan(db *gorm.DB) *gorm.DB {
	return db.Preload("Patient").Preload("Food")
}
