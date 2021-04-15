package diet_plan

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func DietPlanUpdate(dietPlan *model.DietPlan)  (*model.DietPlan,error){
	db :=env.RDB
	diet := model.DietPlan{}
	diet.PreloadDietPlan(db).Model(model.DietPlan{}).First(&diet,diet.ID)
	dietPlan.Patient = diet.Patient
	err := db.Save(dietPlan).Error
	if err!=nil{
		return nil,err
	}else{
		return dietPlan,nil
	}
}