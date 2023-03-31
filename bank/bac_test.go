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
	"testing"

	testfiles "github.com/JuanVF/personal_bot/bank/test_files"
	"github.com/JuanVF/personal_bot/repositories"
)

func TestBankBAC_GetPaymentData(t *testing.T) {
	tName := "TestBankBAC - it_can_detect_payment_data"
	tests := []struct {
		in   string
		want PaymentData
	}{
		{
			in: testfiles.TEST_1_BAC,
			want: PaymentData{
				Body: "other",
				Currency: &repositories.Currency{
					Name: "CRC",
				},
				Amount: 10120.00,
			},
		},
		{
			in: testfiles.TEST_2_BAC,
			want: PaymentData{
				Body: "LA SANWUCHERA TEC2",
				Currency: &repositories.Currency{
					Name: "CRC",
				},
				Amount: 120.00,
			},
		},
		{
			in: testfiles.TEST_3_BAC,
			want: PaymentData{
				Body: "LA SANWUCHERA TEC",
				Currency: &repositories.Currency{
					Name: "USD",
				},
				Amount: 10120.00,
			},
		},
	}

	bac := &BAC{}

	for _, test := range tests {
		got := bac.getPayment(test.in)

		if got.Amount != test.want.Amount {
			logger.TestError(tName, test.want.Amount, got.Amount, t)
		}
		if got.Currency.Name != test.want.Currency.Name {
			logger.TestError(tName, test.want.Currency.Name, got.Currency.Name, t)
		}
		if got.Body != test.want.Body {
			logger.TestError(tName, test.want.Body, got.Body, t)
		}
	}
}
