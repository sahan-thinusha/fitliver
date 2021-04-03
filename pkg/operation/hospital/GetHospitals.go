package hospital

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetHospitals()  ([]*model.Hospital,error){
	db :=env.RDB
	hospital := []*model.Hospital{}
	err := db.Find(&hospital).Error
	if err!=nil{
		return nil,err
	}else{
		return hospital,nil
	}
}
