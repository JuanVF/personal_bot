package bank

import (
	"testing"

	testfiles "github.com/JuanVF/personal_bot/bank/test_files"
	"github.com/JuanVF/personal_bot/repositories"
)

func TestBankBN_GetPaymentData(t *testing.T) {
	tName := "TestBankBN - it_can_detect_payment_data"
	tests := []struct {
		in   string
		want PaymentData
	}{
		{
			in: testfiles.BN_TEST_1,
			want: PaymentData{
				Body: " realizada en *UBER EATS SAN JOSE CRI* ",
				Currency: &repositories.Currency{
					Name: "CRC",
				},
				Amount: 5242,
			},
		},
		{
			in: testfiles.BN_TEST_2,
			want: PaymentData{
				Body: " realizada en *XSOLLA TWITCH LIMASSOL CYP* ",
				Currency: &repositories.Currency{
					Name: "USD",
				},
				Amount: 3.56,
			},
		},
	}

	bn := &BN{}

	for _, test := range tests {
		got := bn.getPayment(test.in)

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
