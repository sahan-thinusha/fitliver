package model


import "github.com/jinzhu/gorm"

type WorkoutPlan struct {
	Model
	PlanName string `json:"plan_name"`
	Patient               *Patient `gorm:"foreignkey:patientID" json:"patient"`
	PatientID             int64
	Workout  []*Workout `gorm:"many2many:workout_workout_plan" json:"workout"`

}
func (WorkoutPlan) TableName() string {
	return "workout_plan"
}
func (m *WorkoutPlan) PreloadWorkoutPlan(db *gorm.DB) *gorm.DB {
	return db.Preload("Patient").Preload("Workout")
}
