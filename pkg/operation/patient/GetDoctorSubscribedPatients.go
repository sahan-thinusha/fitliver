package patient

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"time"
)


func GetDoctorPatients(email string)  ([]*model.Patient,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)

	cs := model.Patient_Consult{}
	cons := []*model.Patient_Consult{}

	cs.PreloadPatient_Consult(db).Joins("join consultations_service on consultations_service.id = patient_consult.consultations_service_id").Joins("join doctor on doctor.id = consultations_service.doctor_id").Where("doctor.id = ?",user.Doctor.ID).Find(&cons)

	patients := []*model.Patient{}
	for _,consalt := range cons {
		t := time.Now()
		diff := t .Sub(consalt.PurchasedAt).Hours() / 24
		if (diff - consalt.Duration) > 0{
			patients = append(patients,consalt.Patient)
		}
	}
	return patients,nil
}
