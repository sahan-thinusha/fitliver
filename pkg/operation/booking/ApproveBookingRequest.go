package booking

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/refund"
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
		stripe.Key = env.StripeSecretKey
		params := &stripe.RefundParams{
			Charge: stripe.String(payment.PaymentId),
		}
		refund.New(params)
		db.Model(&model.BookingPayment{}).Where("id = ?", payment.ID).Update("status","REFUNDED")
	}else if strings.EqualFold(status,env.STATUS_APPROVED) {
			pb := model.Patient_Booking{}
			pb.BookingService = request.Booking
			pb.Hospital = request.Hospital
			pb.Patient = request.Patient
			pb.TimeFrom = request.TimeFrom
			pb.TimeTo= request.TimeTo
			pb.BookedDate = request.BookedDate
			pb.PurchasedAt = request.CreatedAt
			db.Save(&pb)

	}
	return &request,nil

}