package model

import "github.com/jinzhu/gorm"


type Workout struct {
	Model
	Name string `json:"name"`
	VideoLink string `gorm:"size:256"`
	Reps string `json:"reps"`
	WorkoutPlan  []*WorkoutPlan `gorm:"many2many:workout_workout_plan" json:"workout_plan"`

}
func (Workout) TableName() string {
	return "workout"
}
func (m *Workout) PreloadWorkout(db *gorm.DB) *gorm.DB {
	return db
}
