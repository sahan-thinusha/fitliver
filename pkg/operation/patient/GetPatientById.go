package patient

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetPatient(id int)(*model.Patient,error){
	db :=env.RDB
	patient := model.Patient{}
	err := db.First(&patient,id).Error
	if err!=nil{
		return nil,err
	}else{
		return &patient,nil
	}
}
