package repositories

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// Define test data
var paymentTest *CreatePayment = &CreatePayment{
	Amount:      10.5,
	UserId:      1,
	CurrencyId:  1,
	DolarPrice:  3.5,
	Tags:        []string{},
	LastUpdated: "2022-03-27",
	GmailId:     "12345",
	Description: "test description",
}

// Define test data
var paymentTest2 *Payment = &Payment{
	Id:          1,
	Amount:      10.5,
	LastUpdated: "2022-05-03T15:30:00Z",
	Currency:    "USD",
	DolarPrice:  1.0,
	Tags:        []string{"food", "groceries"},
	GmailId:     nil,
	Description: nil,
}

func TestInsertPayment(t *testing.T) {
	tName := "Payment Repository - Insert Payment"

	// Define expected query result
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	// Create a new sqlmock instance
	mock := *db.GetMock()

	// Expect the INSERT query with the test data arguments
	mock.ExpectQuery(`INSERT INTO personal_bot.t_payments(
                            amount, last_updated, user_id, currency_id, dolar_price, tags, gmail_id, description)
                        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
                        RETURNING id`).
		WithArgs(paymentTest.Amount, paymentTest.LastUpdated, paymentTest.UserId, paymentTest.CurrencyId, paymentTest.DolarPrice, `[]`, paymentTest.GmailId, paymentTest.Description).
		WillReturnRows(rows)

	// Call the function to be tested
	paymentId, _ := InsertPayment(paymentTest)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	// Check the function result
	if paymentId != 1 {
		logger.TestError(tName, "Payment created with id: 1", fmt.Sprintf("Payment created with id: %d", paymentId), t)
	}
}

func TestGetPaymentsByUserId(t *testing.T) {
	tName := "Payment Repository - Get Payments By User Id"

	mock := *db.GetMock()

	// We prepare the expected query result for the user
	userRows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated"}).
		AddRow(1, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00")

	mock.ExpectQuery(`SELECT id, name, last_name, google_me, last_updated FROM personal_bot.t_users WHERE id = $1`).
		WithArgs(1).
		WillReturnRows(userRows)

	// We prepare the expected query result for the payments
	rows := sqlmock.NewRows([]string{"payment_id", "amount", "last_updated", "dolar_price", "tags", "currency", "gmail_id", "description"}).
		AddRow(paymentTest2.Id, paymentTest2.Amount, paymentTest2.LastUpdated, paymentTest2.DolarPrice, `["food","groceries"]`, paymentTest2.Currency, paymentTest2.GmailId, paymentTest2.Description)

	mock.ExpectQuery(`SELECT 
                        pay.id payment_id, 
                        pay.amount, 
                        pay.last_updated AT TIME ZONE 'UTC-6',
                        pay.dolar_price,
                        pay.tags,
                        curr.name currency,
                        pay.gmail_id,
                        pay.description
                    FROM personal_bot.t_payments pay
                    INNER JOIN personal_bot.t_currencies curr
                        ON pay.currency_id = curr.id
                    WHERE pay.user_id = $1
                    ORDER BY pay.last_updated DESC`).
		WithArgs(1).
		WillReturnRows(rows)

	// Call function to be tested
	data, _ := GetPaymentsByUserId(1)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if data.Payments[0].Id != 1 {
		logger.TestError(tName, "Payment with id: 1", fmt.Sprintf("Payment with id: %d", data.Payments[0].Id), t)
	}

	if data.Payments[0].Amount != paymentTest2.Amount {
		logger.TestError(tName, fmt.Sprintf("Amount: %f", paymentTest2.Amount), fmt.Sprintf("Amount: %f", data.Payments[0].Amount), t)
	}
}
