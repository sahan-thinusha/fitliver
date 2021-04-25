package booking

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"time"
)

func BookingsForDoctorByDate(date *time.Time,email string)  ([]*model.Patient_Booking,error){
	db :=env.RDB
	dStart := date
	dEnd := date.AddDate(0,0,1)

	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	patientBooking := model.Patient_Booking{}
	patientBookings := []*model.Patient_Booking{}
	err := patientBooking.PreloadPatient_Booking(db).Model(model.Patient_Booking{}).Joins("left join booking_service on patient_booking.booking_service_id = booking_service.id").Joins("left join doctor on booking_service.doctor_id = doctor.id").Where("doctor.id = ? AND booked_date BETWEEN ? AND ?",user.Doctor.ID,dStart,dEnd).Find(&patientBookings).Error
	if err!=nil{
		return nil,err
	}else{
		return patientBookings,nil
	}
}