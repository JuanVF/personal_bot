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
	BNPattern := regexp.MustCompile("(el comprobante de \\*COMPRA\\*)(.*)(el)")
	BNAmountPattern := regexp.MustCompile("(TOTAL)(.*)")
	BNPatternRemove := regexp.MustCompile("(el comprobante de \\*COMPRA\\*)|(el)")
	BNAmountRemove := regexp.MustCompile("(TOTAL:)|CRC|USD")

	data := string(BNPattern.Find([]byte(body)))
	amount := string(BNAmountPattern.Find([]byte(body)))

	if data == "" {
		return nil
	}

	currency := string(CurrenciesRegex.Find([]byte(amount)))

	data = BNPatternRemove.ReplaceAllString(data, "")
	data = BlankSpaces.ReplaceAllString(data, " ")

	amount = BNAmountRemove.ReplaceAllString(amount, "")
	amount = BlankSpace.ReplaceAllString(amount, "")

	amount = strings.ReplaceAll(amount, ",", "")

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
