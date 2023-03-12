package bank

import "github.com/JuanVF/personal_bot/repositories"

type PaymentData struct {
	Body     string
	Currency *repositories.Currency
	Amount   float64
}

type BankHandler interface {
	getCRCValue() (float64, error)
}

type BankMatcher interface {
	getPayment(body string) *PaymentData
}
