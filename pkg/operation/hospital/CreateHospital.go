package hospital

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func AddHospital(hospital *model.Hospital)  (*model.Hospital,error){
	db :=env.RDB
	err := db.Create(&hospital).Error
	if err!=nil{
		return nil,err
	}else{
		return hospital,nil
	}
}
