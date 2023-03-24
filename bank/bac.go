package bank

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/JuanVF/personal_bot/repositories"
)

type BAC struct {
}

func extractEstablishmentName(str string) (string, error) {
	// Define a regular expression to match the establishment name pattern
	re := regexp.MustCompile(`(?s)<p>Comercio:<\/p>\s*</td>\s*<td[^>]*>\s*<p[^>]*>([^<]+)<\/p>`)

	// Find the first match of the regular expression in the string
	match := re.FindStringSubmatch(str)

	if len(match) != 2 {
		return "", fmt.Errorf("invalid input: %s", str)
	}

	return strings.TrimSpace(match[1]), nil
}

// It will convert from one currency to another using BN Data
func (bac BAC) getPayment(body string) *PaymentData {
	establishment, err := extractEstablishmentName(body)

	if err != nil {
		return nil
	}

	// Define a regular expression to match the currency and number pattern
	re := regexp.MustCompile(`([A-Z]{3})\s*([\d,]+\.\d{2})`)

	// Find the first match of the regular expression in the string
	match := re.FindStringSubmatch(body)

	if len(match) != 3 {
		return nil
	}

	// Parse the number to a float64
	num, err := strconv.ParseFloat(strings.Replace(match[2], ",", "", -1), 64)
	if err != nil {
		return nil
	}

	return &PaymentData{
		Body: establishment,
		Currency: &repositories.Currency{
			Name: match[1],
		},
		Amount: num,
	}
}
