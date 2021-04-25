package consultation_service



import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"fmt"
	"strings"
)

func ConsultationServiceDelete(id int64,email string,role string)  (*model.ConsultationService,error){
	db :=env.RDB
	db.LogMode(true)
	user := model.User{}
	fmt.Println("GG")
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	consultationService := model.ConsultationService{}
	consultationService.PreloadConsultationServiceForUser(db).First(&consultationService,id)


	if consultationService.DoctorID == user.Doctor.ID || strings.EqualFold(env.ADMIN,role){
		err := db.Delete(&consultationService).Error
		if err!=nil{
			return nil,err
		}else{
			return &consultationService,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}