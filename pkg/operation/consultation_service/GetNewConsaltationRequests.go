package consultation_service

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func NewConsultationRequestsForDoctor(email string)  ([]*model.ConsultationRequest,error){
	db :=env.RDB
	db.LogMode(true)
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	consultationRequest := model.ConsultationRequest{}
	consultationRequests := []*model.ConsultationRequest{}
	err := consultationRequest.PreloadConsultationRequest(db).Model(model.ConsultationRequest{}).Joins("left join package on package.id = consultation_request.package_id").Joins("left join consultation_service on consultation_service.id = package.consultationservice_id").Where("consultationservice.doctor_id = ? AND consultation_request.status = ?",user.Doctor.ID,env.STATUS_NEW).Find(consultationRequests).Error
	if err!=nil{
		return nil,err
	}else{
		return consultationRequests,nil
	}
}