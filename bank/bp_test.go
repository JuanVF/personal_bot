package bank

import (
	"testing"

	testfiles "github.com/JuanVF/personal_bot/bank/test_files"
	"github.com/JuanVF/personal_bot/repositories"
)

func TestBankBP_it_can_detect_payment_data(t *testing.T) {
	tName := "TestBankBP - it_can_detect_payment_data"
	tests := []struct {
		in   string
		want PaymentData
	}{
		{
			in: testfiles.BP_TEST_1,
			want: PaymentData{
				Body: " APPLE.COM/BILL xxx-xxx-xxxx US por",
				Currency: &repositories.Currency{
					Name: "USD",
				},
				Amount: 1.99,
			},
		},
		{
			in: testfiles.BP_TEST_2,
			want: PaymentData{
				Body: " DL DIDI FOODS SAN JOSE CR por",
				Currency: &repositories.Currency{
					Name: "CRC",
				},
				Amount: 4130,
			},
		},
	}

	bp := &BP{}

	for _, test := range tests {
		got := bp.getPayment(test.in)

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
