package doctor

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func DoctorRegister(doctor *model.Doctor)  (*model.Doctor,error){
	db :=env.RDB
	err := db.Create(doctor).Error

	if err!=nil{
		return nil,err
	}else{
		return doctor,nil
	}
}