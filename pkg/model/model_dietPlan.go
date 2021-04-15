package model

import "github.com/jinzhu/gorm"

type  DietPlan struct{
	Model
	Food      []*FoodData `json:"food"`
	Type  string `json:"type"`
	MealName string `json:"mealName"`
	Patient     *Patient `gorm:"foreignkey:patientID" json:"patient"`
	PatientID int64
}

func (DietPlan) TableName() string {
	return "diet_plan"
}
func (m *DietPlan) PreloadDietPlan(db *gorm.DB) *gorm.DB {
	return db.Preload("FoodData").Preload("Patient")
}
