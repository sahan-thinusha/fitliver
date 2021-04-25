package booking

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func NewBookingRequestsForDoctor(email string)  ([]*model.BookingRequest,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	doc := user.Doctor
	bookingRequest := model.BookingRequest{}
	bookingRequests := []*model.BookingRequest{}
	bookingRequest.PreloadBookingRequest(db).Model(model.BookingRequest{}).Joins("left join booking_service on booking_service.id = booking_request.booking_id").Joins("left join doctor on doctor.id = booking_service.doctor_id").Where("doctor.id = ? AND booking_request.status = ?",doc.ID,env.STATUS_NEW).Find(&bookingRequests)
		return bookingRequests,nil
}