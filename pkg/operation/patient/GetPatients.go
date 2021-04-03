package patient

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetPatients()  ([]*model.Patient,error){
	db :=env.RDB
	patients := []*model.Patient{}
	err := db.Find(&patients).Error
	if err!=nil{
		return nil,err
	}else{
		return patients,nil
	}
}
