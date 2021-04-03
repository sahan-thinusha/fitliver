package doctor

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetDoctors()  ([]*model.Doctor,error){
	db :=env.RDB
	doctors := []*model.Doctor{}
	doctor := model.Doctor{}
	err := doctor.PreloadDoctor(db).Find(&doctors).Error
	if err!=nil{
		return nil,err
	}else{
		return doctors,nil
	}
}
