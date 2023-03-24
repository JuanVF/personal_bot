package bank

import (
	"regexp"
	"strconv"

	"github.com/JuanVF/personal_bot/repositories"
)

type BP struct {
}

// It will convert from one currency to another using BP Data
func (bp BP) getPayment(body string) *PaymentData {
	data := bp.getData(body)
	amount := bp.getAmount(body)

	if data == "" {
		return nil
	}

	currency := bp.getCurrency(body)

	return &PaymentData{
		Body: data,
		Currency: &repositories.Currency{
			Name: currency,
		},
		Amount: amount,
	}
}

// getData extracts and cleans the data from the input body string
// using regular expressions.
// It returns the cleaned data string.
func (bp BP) getData(body string) string {
	BPPattern := regexp.MustCompile("(fue utilizada en)(.*)(por)")
	BPPatternRemove := regexp.MustCompile("(fue utilizada en)|(Si no reconoce)")

	data := string(BPPattern.Find([]byte(body)))

	data = BlankSpaces.ReplaceAllString(data, " ")
	data = BPPatternRemove.ReplaceAllString(data, "")

	return data
}

// getCurrency extracts the currency from the provided amount string following BP format.
// It looks for the currency code in the amount string and uses the
// CurrenciesRegex regular expression to extract the currency name.
// If no currency is found, an empty string is returned.
func (bp BP) getCurrency(body string) string {
	BPAmount := regexp.MustCompile("(\\d)+\\.(\\d)+\\s(USD|CRC)")
	amount := string(BPAmount.Find([]byte(body)))

	currency := string(CurrenciesRegex.Find([]byte(amount)))

	return currency
}

// getAmount extracts the amount from the provided body string following BP format,
// removes the unnecessary text and symbols, and returns the numeric amount as float64.
// If there's an error parsing the amount, it returns 0.
func (bp BP) getAmount(body string) float64 {
	BPAmount := regexp.MustCompile("(\\d)+\\.(\\d)+\\s(USD|CRC)")
	BPAmountRemove := regexp.MustCompile("USD|CRC")

	amount := string(BPAmount.Find([]byte(body)))

	amount = BPAmountRemove.ReplaceAllString(amount, "")
	amount = BlankSpace.ReplaceAllString(amount, "")

	numericAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return 0
	}

	return numericAmount
}
