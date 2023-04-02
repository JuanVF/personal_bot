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
	IdBank:      1,
}

// Define test data
var paymentTest2 *Payment = &Payment{
	Id:          1,
	Amount:      10.5,
	LastUpdated: "2022-05-03T15:30:00Z",
	Currency:    "USD",
	DolarPrice:  1.0,
	Tags:        []string{"food", "groceries"},
	GmailId:     &(&struct{ d string }{d: "abc"}).d,
	Description: &(&struct{ d string }{d: "Hi"}).d,
	IdBank:      &(&Bank{Id: 1}).Id,
	Bank:        &Bank{Id: 1, Name: "Bank of America"},
}

func TestInsertPayment(t *testing.T) {
	tName := "Payment Repository - Insert Payment"

	// Define expected query result
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	// Create a new sqlmock instance
	mock := *db.GetMock()

	// Expect the INSERT query with the test data arguments
	mock.ExpectQuery(`INSERT INTO personal_bot.t_payments(
						amount, last_updated, user_id, currency_id, dolar_price, tags, gmail_id, description, id_bank)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
					RETURNING id`).
		WithArgs(paymentTest.Amount, paymentTest.LastUpdated, paymentTest.UserId, paymentTest.CurrencyId, paymentTest.DolarPrice, `[]`, paymentTest.GmailId, paymentTest.Description, paymentTest.IdBank).
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
	userRows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated", "weight", "height", "activity_level_id", "name", "description"}).
		AddRow(1, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00", userTest.Weight, userTest.Height, userTest.ActivityLevelId, "Sedentary", "Little or no exercise")

	mock.ExpectQuery(`SELECT 
			u.id, u.name, u.last_name, u.google_me, u.last_updated, u.weight, u.height, u.activity_level_id,
			al.name, al.description
		FROM personal_bot.t_users u
		INNER JOIN personal_bot.t_activity_levels al ON al.id = u.activity_level_id
		WHERE u.id = $1`).
		WithArgs(1).
		WillReturnRows(userRows)
	// We prepare the expected query result for the payments
	rows := sqlmock.NewRows([]string{"payment_id", "amount", "last_updated", "dolar_price", "tags", "currency", "gmail_id", "description", "id_bank", "bank_name"}).
		AddRow(paymentTest2.Id, paymentTest2.Amount, paymentTest2.LastUpdated, paymentTest2.DolarPrice, `["food","groceries"]`, paymentTest2.Currency, paymentTest2.GmailId, paymentTest2.Description, paymentTest2.Bank.Id, paymentTest2.Bank.Name)

	mock.ExpectQuery(`SELECT 
						pay.id payment_id, 
						pay.amount, 
						pay.last_updated AT TIME ZONE 'UTC-6',
						pay.dolar_price,
						pay.tags,
						curr.name currency,
						pay.gmail_id,
						pay.description,
						pay.id_bank,
						bank.name bank_name
					FROM personal_bot.t_payments pay
					INNER JOIN personal_bot.t_currencies curr
						ON pay.currency_id = curr.id
					INNER JOIN personal_bot.t_banks bank
						ON pay.id_bank = bank.id
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

	if data.Payments[0] == nil {
		logger.TestError(tName, "Payment created", "Payment not created", t)
	}

	if data.Payments[0].Id != 1 {
		logger.TestError(tName, "Payment with id: 1", fmt.Sprintf("Payment with id: %d", data.Payments[0].Id), t)
	}

	if data.Payments[0].Amount != paymentTest2.Amount {
		logger.TestError(tName, fmt.Sprintf("Amount: %f", paymentTest2.Amount), fmt.Sprintf("Amount: %f", data.Payments[0].Amount), t)
	}
}
