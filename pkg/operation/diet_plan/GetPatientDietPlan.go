package diet_plan


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetPatientDietPlans(id int)([]*model.DietPlan,error){
	db :=env.RDB
	dietPlan := model.DietPlan{}
	var dietPlans []*model.DietPlan
	err := dietPlan.PreloadDietPlan(db).Where("patient_id = ?",id).Find(&dietPlans).Error
	if err!=nil{
		return nil,err
	}else{
		return dietPlans,nil
	}
}
