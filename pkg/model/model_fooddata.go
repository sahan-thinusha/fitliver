package model

import "github.com/jinzhu/gorm"

type   FoodData struct{
	Label     string `json:"label"`
	Nutrients struct {
		EnercKcal float64     `json:"enerc_kcal"`
		Procnt    float64 `json:"procnt"`
		Fat       float64 `json:"fat"`
		Chocdf    float64     `json:"chocdf"`
		Fibtg     float64     `json:"fibtg"`
	} `json:"nutrients"`
	Category      string `json:"category"`
	Categorylabel string `json:"categoryLabel"`
	DietPlan     DietPlan `gorm:"foreignkey:dietPlanID" json:"dietPlan"`
	DietPlanID int64
}

func (FoodData) TableName() string {
	return "food_data"
}
func (m *FoodData) PreloadFoodData(db *gorm.DB) *gorm.DB {
	return db
}
