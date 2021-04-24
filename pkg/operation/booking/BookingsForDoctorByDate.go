package booking

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BookingsForDoctorByDate(date string,email string)  ([]*model.Patient_Booking,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	patientBooking := model.Patient_Booking{}
	patientBookings := []*model.Patient_Booking{}
	err := patientBooking.PreloadPatient_Booking(db).Model(model.Patient_Booking{}).Joins("joins booking_service on patient_booking.booking_service_id = booking_service.id").Joins("joins doctor on booking_service.doctor_id = doctor.id").Where("doctor.id = ? AND booked_date = ?",user.Doctor.ID,date).Find(patientBookings).Error
	if err!=nil{
		return nil,err
	}else{
		return patientBookings,nil
	}
}