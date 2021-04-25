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
	doctorList := []*model.Doctor{}
	for _,doctor := range doctors{
		summary := RateSummary{}
		db.Model(model.Rating{}).Select("AVG(rate) AS rate").Where("doctor_id = ?",doctor.ID).Scan(&summary)
		doctor.Rate= summary.Rate
		doctorList = append(doctorList,doctor)
	}

	if err!=nil{
		return nil,err
	}else{
		return doctorList,nil
	}
}
