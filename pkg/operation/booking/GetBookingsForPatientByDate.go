package booking


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"time"
)

func BookingsForPatientByDate(date *time.Time,email string)  ([]*model.Patient_Booking,error){
	db :=env.RDB
	user := model.User{}
	dStart := date
	dEnd := date.AddDate(0,0,1)

	user.PreloadPatient(db).Model(model.User{}).Where("email = ?",email).First(&user)
	patientBooking := model.Patient_Booking{}
	patientBookings := []*model.Patient_Booking{}
	err := patientBooking.PreloadPatient_Booking(db).Model(model.Patient_Booking{}).Where("patient_id = ? AND booked_date BETWEEN ? AND ?",user.Patient.ID,dStart,dEnd).Find(&patientBookings).Error
	if err!=nil{
		return nil,err
	}else{
		return patientBookings,nil
	}
}