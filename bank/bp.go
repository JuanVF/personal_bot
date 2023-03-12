package bank

import (
	"regexp"
	"strconv"

	"github.com/JuanVF/personal_bot/repositories"
)

type BP struct {
}

// It will convert from one currency to another using BN Data
func (bp BP) getPayment(body string) *PaymentData {
	BPPattern := regexp.MustCompile("(fue utilizada en)(.*)(por)")
	BPAmount := regexp.MustCompile("(\\d)+\\.(\\d)+\\s(USD|CRC)")
	BPPatternRemove := regexp.MustCompile("(fue utilizada en)|(Si no reconoce)")

	data := string(BPPattern.Find([]byte(body)))
	amount := string(BPAmount.Find([]byte(body)))

	currency := string(CurrenciesRegex.Find([]byte(amount)))

	data = BlankSpaces.ReplaceAllString(data, " ")
	data = BPPatternRemove.ReplaceAllString(data, "")

	amount = CurrenciesRegex.ReplaceAllString(amount, "")
	amount = BlankSpace.ReplaceAllString(amount, "")

	numericAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return nil
	}

	return &PaymentData{
		Body: data,
		Currency: &repositories.Currency{
			Name: currency,
		},
		Amount: numericAmount,
	}
}
