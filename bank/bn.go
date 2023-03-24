package bank

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/JuanVF/personal_bot/repositories"
)

type BN struct {
}

// It will convert from one currency to another using BN Data
func (bn BN) getPayment(body string) *PaymentData {
	data := bn.getData(body)
	amount := bn.getAmount(body)

	if data == "" {
		return nil
	}

	currency := bn.getCurrency(body)

	return &PaymentData{
		Body: data,
		Currency: &repositories.Currency{
			Name: currency,
		},
		Amount: amount,
	}
}

// getCurrency extracts the currency from the given BN payment data.
// It looks for the string "TOTAL" in the payment data and uses the
// CurrenciesRegex regular expression to extract the currency name.
// If no currency is found, an empty string is returned.
func (bn BN) getCurrency(body string) string {
	BNAmountPattern := regexp.MustCompile("(TOTAL)(.*)")

	currency := string(CurrenciesRegex.Find([]byte(BNAmountPattern.Find([]byte(body)))))

	return currency
}

// getAmount extracts the amount from the provided body string following BN format,
// removes the unnecessary text and symbols, and returns the numeric amount as float64.
// If there's an error parsing the amount, it returns 0.
func (bn BN) getAmount(body string) float64 {
	BNAmountPattern := regexp.MustCompile("(TOTAL)(.*)")
	BNAmountRemove := regexp.MustCompile("(TOTAL:)|CRC|USD")

	amount := string(BNAmountPattern.Find([]byte(body)))

	amount = BNAmountRemove.ReplaceAllString(amount, "")
	amount = BlankSpace.ReplaceAllString(amount, "")

	amount = strings.ReplaceAll(amount, ",", "")

	numericAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return 0
	}

	return numericAmount
}

// getData extracts and cleans the data from the input body string
// using regular expressions.
// It returns the cleaned data string.
func (bn BN) getData(body string) string {
	BNPattern := regexp.MustCompile("(el comprobante de \\*COMPRA\\*)(.*)(el)")
	BNPatternRemove := regexp.MustCompile("(el comprobante de \\*COMPRA\\*)|(el)")

	data := string(BNPattern.Find([]byte(body)))

	data = BNPatternRemove.ReplaceAllString(data, "")
	data = BlankSpaces.ReplaceAllString(data, " ")

	return data
}
