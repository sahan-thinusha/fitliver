package booking

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
)

func ApproveBookingRequest(id int64,status string,email string)  (*model.BookingRequest,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)

	err := db.Model(&model.BookingRequest{}).Where("id = ?", id).Update("status",status).Error
	if err!=nil {
		return nil, err
	}
	request := model.BookingRequest{}
	request.PreloadBookingRequest(db).First(&request,id)
	if strings.EqualFold(status,env.STATUS_REJECTED) {
		payment := model.BookingPayment{}
		db.Model(&model.BookingPayment{}).Where("booking_request_id = ?",request.Booking.ID).First(&payment)
		db.Model(&model.BookingPayment{}).Where("id = ?", payment.ID).Update("status","REFUNDED")
	}
	return &request,nil

}