package doctor


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func AddRating(rating *model.Rating,email string)  (*model.Rating,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	rating.User = &user
	doctor := model.Doctor{}
	db.Model(model.Doctor{}).First(&doctor,rating.DoctorID)
	rating.Doctor = &doctor
	err := db.Create(&rating).Error
	if err!=nil{
		return nil,err
	}else{
		return rating,nil
	}
}
