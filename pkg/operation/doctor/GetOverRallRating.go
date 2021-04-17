package doctor


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

type RateSummary struct{
	avgRate float64 `gorm:"column:avgRate"`
}

func GetOverallRatingForDoctor(id int64)  (*model.Doctor,error){
	db :=env.RDB
	summary := RateSummary{}
	db.Model(model.Rating{}).Select("AVG(rating) AS avgRate").Where("doctor = ?",id).Scan(&summary)
	doctor := model.Doctor{}
	db.Model(model.Doctor{}).First(&doctor,id)
	doctor.Rate = summary.avgRate
	return &doctor,nil
}