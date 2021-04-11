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
	pack := model.Package{}
	pack.PreloadPackage(db).Model(model.Package{}).First(&pack, payment.Package.ID)
	payment.ReceiptEmail = email
	payment.Amount = pack.Amount
	payment.Name = pack.Name + " " + pack.ConsultationService.Name
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
	db.Save(payment)
	return payment, nil
}
