package consultation_service



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ConsultationServiceForDoctor(email string)  ([]*model.ConsultationService,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	consultationService := model.ConsultationService{}
	consultationServices := []*model.ConsultationService{}
	err := consultationService.PreloadConsultationService(db).Model(model.ConsultationService{}).Where("doctor_id = ?",user.Doctor.ID).Find(consultationServices).Error
	if err!=nil{
		return nil,err
	}else{
		return consultationServices,nil
	}
}