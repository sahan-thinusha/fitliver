package consultation_service



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func ConsultationServicesGetAll()  ([]*model.ConsultationService,error){
	db :=env.RDB
	consultationService := model.ConsultationService{}
	consultationServices := []*model.ConsultationService{}
	err := consultationService.PreloadConsultationService(db).Model(model.ConsultationService{}).Find(consultationServices).Error
	if err!=nil{
		return nil,err
	}else{
		return consultationServices,nil
	}
}