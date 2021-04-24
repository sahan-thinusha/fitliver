package booking


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BookingsForPatientByDate(date string,email string)  ([]*model.Patient_Booking,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadPatient(db).Model(model.User{}).Where("email = ?",email).First(&user)
	patientBooking := model.Patient_Booking{}
	patientBookings := []*model.Patient_Booking{}
	err := patientBooking.PreloadPatient_Booking(db).Model(model.Patient_Booking{}).Where("patient_id = ? AND booked_date = ?",user.Patient.ID,date).Find(patientBookings).Error
	if err!=nil{
		return nil,err
	}else{
		return patientBookings,nil
	}
}