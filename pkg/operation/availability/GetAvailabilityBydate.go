package availability

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetAvailabilityByDate(date string,email string)  ([]*model.Availability,error){
	db :=env.RDB
	user := model.User{}

	db.Model(model.User{}).Where("email = ?",email).First(&user)
	avalabilities := []*model.Availability{}
	availability := model.Availability{}
	availability.PreloadAvailability(db).Where("doctor_id = ? AND available_date = ?",user.Doctor.ID,date).Find(&avalabilities)
	return avalabilities,nil
}