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

This file was created by Gerald Zamora (geraldzmt@gmail.com).
*/
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
	re := regexp.MustCompile(`(?s)<p>Comercio:<\/p>\s*<\/td>\s*<td[^>]*>\s*<p[^>]*>([^<]+)<\/p>`)

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
		Amount:   num,
		BankName: "BAC",
	}
}
