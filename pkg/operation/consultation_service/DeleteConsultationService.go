package consultation_service



import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
)

func ConsultationServiceDelete(id int64,email string,role string)  (*model.ConsultationService,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	consultationService := model.ConsultationService{}
	consultationService.PreloadConsultationServiceForUser(db).First(&consultationService,id)
	if consultationService.Doctor.ID == user.Doctor.ID || strings.EqualFold(env.ADMIN,role){
		err := db.Delete(consultationService).Error
		if err!=nil{
			return nil,err
		}else{
			return &consultationService,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}