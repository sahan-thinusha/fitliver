package availability



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func AddAvailabiility(availability *model.Availability,email string)  (*model.Availability,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	availability.Doctor = user.Doctor
	availability.IsFree = true
	hospital := model.Hospital{}
	db.Model(model.Hospital{}).First(&hospital,availability.HospitalID)
	availability.Hospital = &hospital
	err := db.Create(availability).Error
	if err!=nil{
		return nil,err
	}else{
		return availability,nil
	}
}