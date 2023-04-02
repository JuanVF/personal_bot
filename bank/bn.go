/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
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
		Amount:   amount,
		BankName: "BN",
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
