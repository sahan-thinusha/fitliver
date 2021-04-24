package booking


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func NewBookingRequestsForDoctor(email string)  ([]*model.BookingRequest,error){
	db :=env.RDB
	db.LogMode(true)
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	bookingRequest := model.BookingRequest{}
	bookingRequests := []*model.BookingRequest{}
	err := bookingRequest.PreloadBookingRequest(db).Model(model.BookingRequest{}).Joins("left join booking_service on booking_service.id = booking_request.booking_service_id").Joins("left join doctor on doctor.id = booking_service.doctor_id").Where("doctor.id = ? AND booking_request.status = ?",user.Doctor.ID,env.STATUS_NEW).Find(bookingRequests).Error
	if err!=nil{
		return nil,err
	}else{
		return bookingRequests,nil
	}
}