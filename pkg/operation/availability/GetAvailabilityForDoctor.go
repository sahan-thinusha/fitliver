package availability

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetAvailabilityByDateDoctor(date string,id int64)  ([]*model.Availability,error){
	db :=env.RDB
	avalabilities := []*model.Availability{}
	availability := model.Availability{}
	availability.PreloadAvailability(db).Where("doctor_id = ? AND available_date = ?",id,date).Find(&avalabilities)
	return avalabilities,nil
}