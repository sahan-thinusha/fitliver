package healthrecord


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func UpdateHealthRecord(healthRecord *model.HealthRecord,email string)  (*model.HealthRecord,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	healthRecord.Doctor = user.Doctor
	u := model.Patient{}
	u.PreloadPatient(db).Model(model.Patient{}).First(&u,healthRecord.PatientID)
	healthRecord.Patient = &u
	err := db.Save(&healthRecord).Error
	if err!=nil{
		return nil,err
	}else{
		return healthRecord,nil
	}
}
