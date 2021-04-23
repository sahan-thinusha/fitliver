package workout


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func UpdateWorkoutPlan(workoutplan *model.WorkoutPlan,email string)  (*model.WorkoutPlan,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadPatient(db).Model(model.User{}).Where("email = ?",email).First(&user)
	workoutplan.Patient = user.Patient
	err := db.Save(&workoutplan).Error
	if err!=nil{
		return nil,err
	}else{
		return workoutplan,nil
	}
}

