package doctor


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"fmt"
)

type RateSummary struct{
	Rate float64
}

func GetOverallRatingForDoctor(id int64)  (*model.Doctor,error){
	db :=env.RDB
	summary := RateSummary{}
	db.Model(model.Rating{}).Select("AVG(rate) AS rate").Where("doctor_id = ?",id).Scan(&summary)
	fmt.Println(summary.Rate)
	fmt.Println(summary.Rate)
	doctor := model.Doctor{}
	db.Model(model.Doctor{}).First(&doctor,id)
	doctor.Rate = summary.Rate
	return &doctor,nil
}