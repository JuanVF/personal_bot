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