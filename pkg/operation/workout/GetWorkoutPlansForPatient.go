package workout



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetWorkoutPlansForPatient(id int64)  ([]*model.WorkoutPlan,error){
	db :=env.RDB
	workoutplan := model.WorkoutPlan{}
	workoutPlans :=[]*model.WorkoutPlan{}
	workoutplan.PreloadWorkoutPlan(db).Model(model.WorkoutPlan{}).Where("patient_id = ?",id).Find(&workoutPlans)
	return workoutPlans,nil
}