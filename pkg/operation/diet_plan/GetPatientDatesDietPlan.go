package diet_plan

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"time"
)

func GetPatientDietPlansByDate(id int64,starting time.Time,ending time.Time)([]*model.DietPlan,error){
	db :=env.RDB
	dietPlan := model.DietPlan{}
	var dietPlans []*model.DietPlan
	err := dietPlan.PreloadDietPlan(db).Where("patient_id = ? AND created_at BETWEEN ? AND ?",id,starting,ending).Find(&dietPlans).Error
	if err!=nil{
		return nil,err
	}else{
		return dietPlans,nil
	}
}
