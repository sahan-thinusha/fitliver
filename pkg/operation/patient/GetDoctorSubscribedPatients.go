package patient

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"fmt"
	"time"
)


func GetDoctorPatients(email string)  ([]*model.Patient,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)

	cs := model.Patient_Consult{}
	cons := []*model.Patient_Consult{}

	cs.PreloadPatient_Consult(db).Joins("join consultation_service on consultation_service.id = patient_consult.consultation_service_id").Joins("join doctor on doctor.id = consultation_service.doctor_id").Where("doctor.id = ?",user.Doctor.ID).Find(&cons)

	patients := []*model.Patient{}
	patientData := []*model.Patient{}

	for _, consult := range cons {
		fmt.Println(consult.Patient.Name)
		t := time.Now()
		diff := t .Sub(consult.PurchasedAt).Hours() / 24
		fmt.Println(diff)
		if (consult.Duration - diff) > 0{
			patients = append(patients, consult.Patient)
		}
	}
	check := make(map[*model.Patient]int)

	for _, val := range patients {
		check[val] = 1
	}

	for data, _ := range check {
		patientData = append(patientData,data)
	}
	return patientData,nil
}
