package diet_plan

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func DietPlanSave(dietPlan *model.DietPlan,email string)  (*model.DietPlan,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadPatient(db).Model(model.User{}).Where("email = ?",email).First(&user)
	dietPlan.Patient = user.Patient
	err := db.Create(dietPlan).Error
	if err!=nil{
		return nil,err
	}else{
		return dietPlan,nil
	}
}