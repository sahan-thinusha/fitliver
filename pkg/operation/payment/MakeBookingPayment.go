package payment



import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func ProceedBookingPayment(payment *model.BookingPayment, token string, email string) (*model.BookingPayment, error) {
	stripe.Key = env.StripeSecretKey
	db := env.RDB
	user := model.User{}
	user.PreloadPatient(db).Model(model.User{}).Where("email = ?", email).First(&user)
	payment.ReceiptEmail = email
	booking := model.BookingService{}
	booking.PreloadBookingService(db).Model(model.Package{}).First(&booking, payment.Booking.ID)
	payment.ReceiptEmail = email
	payment.Amount = booking.Amount
	payment.Name = booking.Name
	name := ""
	if user.Patient != nil {
		name = user.Patient.Name
	}
	stripeCharge, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount * 100),
		Currency:     stripe.String(string(stripe.CurrencyLKR)),
		Description:  stripe.String(payment.Name),
		Source:       &stripe.SourceParams{Token: stripe.String(token)},
		ReceiptEmail: stripe.String(payment.ReceiptEmail),
		Customer:     stripe.String(name)})
	if stripeCharge != nil {
		payment.PaymentId = stripeCharge.ID
	}
	if err != nil {
		payment.Status = "Unsuccessful"
		db.Save(payment)
		return nil, errors.New("Payment Unsuccessful...")
	}
	payment.Status = "Successful"
	tx := db.Begin()
	err =tx.Save(payment).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Payment Unsuccessful...")
	}
	bookingRequest := model.BookingRequest{}
	bookingRequest.Status = env.STATUS_NEW
	bookingRequest.Booking = &booking
	bookingRequest.Patient = user.Patient
	err = tx.Create(&bookingRequest).Error

	payment.BookingRequest = bookingRequest
	tx.Save(&payment)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Payment Unsuccessful...")
	}


	tx.Commit()

	return payment, nil
}
