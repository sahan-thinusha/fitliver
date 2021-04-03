package doctor

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strconv"
	"strings"
)

func DoctorRegister(doctor *model.Doctor,hospitals string)  (*model.Doctor,error){
	db :=env.RDB
	hospitalList := []*model.Hospital{}
	if hospitals != ""{
		s := strings.Split(hospitals, ",")
		var ids = []int{}
		for _, i := range s {
			j, err := strconv.Atoi(i)
			if err != nil {
				return nil,err
			}
			ids = append(ids, j)
		}
		db.Model(model.Hospital{}).Find(&hospitalList,ids)
	}
	doctor.Hospital = hospitalList
	err := db.Create(doctor).Error
	if err!=nil{
		return nil,err
	}else{
		return doctor,nil
	}
}