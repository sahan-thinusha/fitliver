package doctor

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetDoctor(id int)  (*model.Doctor,error){
	db :=env.RDB
	doctor := model.Doctor{}
	err := db.Where("id = ?",id).First(&doctor).Error
	if err!=nil{
		return nil,err
	}else{
		return &doctor,nil
	}
}
