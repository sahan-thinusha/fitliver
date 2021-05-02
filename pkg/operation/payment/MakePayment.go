package payment

import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func ProceedPayment(payment *model.Payment, token string, email string) (*model.Payment, error) {
	stripe.Key = env.StripeSecretKey
	db := env.RDB
	user := model.User{}
	user.PreloadPatient(db).Model(model.User{}).Where("email = ?", email).First(&user)
	payment.ReceiptEmail = email
	payment.Patient = user.Patient
	pack := model.Package{}
	db.Model(model.Package{}).First(&pack, payment.PackageID)
	cons := model.ConsultationService{}
	db.Model(model.ConsultationService{}).First(&cons, pack.ConsultationServiceID)
	pack.ConsultationService = &cons
	payment.ReceiptEmail = email
	payment.Amount = pack.Amount
	payment.Name =  pack.ConsultationService.Name + " " + pack.Name

		stripeCharge, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount * 100),
		Currency:     stripe.String(string(stripe.CurrencyLKR)),
		Description:  stripe.String(payment.Name),
		//Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		Source:       &stripe.SourceParams{Token: stripe.String(token)},
		ReceiptEmail: stripe.String(payment.ReceiptEmail)})
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
	consultationRequest := model.ConsultationRequest{}
	consultationRequest.Status = env.STATUS_NEW
	consultationRequest.Package = &pack
	consultationRequest.Patient = user.Patient
	err = tx.Create(&consultationRequest).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Payment Unsuccessful...")
	}


	tx.Commit()

	return payment, nil
}
