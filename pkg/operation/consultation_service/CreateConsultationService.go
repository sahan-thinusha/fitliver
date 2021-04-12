package consultation_service

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ConsultationServiceCreate(consultantService *model.ConsultationService,email string)  (*model.ConsultationService,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	consultantService.Doctor = user.Doctor
	err := db.Create(consultantService).Error
	if err!=nil{
		return nil,err
	}else{
		return consultantService,nil
	}
}