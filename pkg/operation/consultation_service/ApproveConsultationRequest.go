package consultation_service

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
	"time"
)

func ApproveConsultationRequest(id int64,status string,email string)  (*model.ConsultationRequest,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)

	err := db.Model(&model.ConsultationRequest{}).Where("id = ?", id).Update("status",status).Error
	if err!=nil {
		return nil, err
	}
	request := model.ConsultationRequest{}
	request.PreloadConsultationRequest(db).First(&request,id)
	if strings.EqualFold(status,env.STATUS_APPROVED){
		pc := model.Patient_Consult{}
		b := db.Model(&model.Patient_Consult{}).Where("consultations_service_id = ?",request.Package.ConsultationService.ID).First(&pc).RecordNotFound()
		if b{
			patientConsalt := model.Patient_Consult{}
			patientConsalt.ConsultationService = request.Package.ConsultationService
			patientConsalt.Duration = request.Package.Duration
			patientConsalt.Patient = request.Patient
			patientConsalt.PurchasedAt = time.Now()
			db.Create(&patientConsalt)
		}else{

		}
	}
	return &request,nil
}