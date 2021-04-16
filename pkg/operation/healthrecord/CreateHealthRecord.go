package healthrecord


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func AddHealthRecord(healthRecord *model.HealthRecord,email string)  (*model.HealthRecord,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	healthRecord.Doctor = user.Doctor
	u := model.User{}
	u.PreloadPatient(db).Model(model.User{}).First(&u,healthRecord.Patient.ID)
	healthRecord.Patient = u.Patient
	err := db.Create(&healthRecord).Error
	if err!=nil{
		return nil,err
	}else{
		return healthRecord,nil
	}
}
