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
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type CreatePayment struct {
	Amount      float64
	UserId      int
	CurrencyId  int
	DolarPrice  float64
	Tags        []string
	LastUpdated string
	GmailId     string
	Description string
	IdBank      int
}

type Payment struct {
	Id          int
	Amount      float64
	LastUpdated string
	Currency    string
	DolarPrice  float64
	Tags        []string
	GmailId     *string
	Description *string
	IdBank      *int
	Bank        *Bank
}

type PaymentSummary struct {
	Month       string
	TotalCount  int
	TotalAmount float64
	CurrencyId  int
	Currency    *Currency
}

type UserPaymentsSummary struct {
	User            *User
	PaymentsSummary []*PaymentSummary
}

type Bank struct {
	Id   int
	Name string
}

type bankCacheItem struct {
	bank       *Bank
	expiration time.Time
}

type UserPayments struct {
	User     *User
	Payments []*Payment
}

var bankCache = make(map[string]bankCacheItem)

// Inserts a payment and its tags to the database
func InsertPayment(payment *CreatePayment) (int, error) {
	var paymentId int = 0

	statement := `INSERT INTO personal_bot.t_payments(
					amount, last_updated, user_id, currency_id, dolar_price, tags, gmail_id, description, id_bank)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
				RETURNING id`

	tags, err := json.Marshal(payment.Tags)

	if err != nil {
		return 0, err
	}

	err = db.GetConnection().
		QueryRow(statement, payment.Amount, payment.LastUpdated, payment.UserId, payment.CurrencyId, payment.DolarPrice, string(tags), payment.GmailId, payment.Description, payment.IdBank).
		Scan(&paymentId)

	if err != nil {
		logger.Error("Payment Repository - Insert Payment", err.Error())

		return 0, err
	}

	return paymentId, nil
}

// I am optimizing this insert later i am lazy rn
func InsertPayments(payments []*CreatePayment) error {
	for _, p := range payments {
		InsertPayment(p)
	}

	return nil
}

// The GetSummaryOfCertainPaymentsByDatesAndUserId function retrieves a summary of payments based on the provided user ID,
// matching criteria, and date range. It returns a UserPaymentsSummary object.
func GetSummaryOfCertainPaymentsByDatesAndUserId(userId int, matcher string, from string, to string) (*UserPaymentsSummary, error) {
	var payments *UserPaymentsSummary = &UserPaymentsSummary{
		PaymentsSummary: make([]*PaymentSummary, 0),
	}

	user, err := GetUser(userId)

	if err != nil {
		logger.Error("Currencies Repository - Get Payments - Get Users", err.Error())

		return nil, err
	}

	payments.User = user

	statement := `SELECT DATE_TRUNC('month', p.last_updated) AS month,
					COUNT(p.id) AS total_count,
					SUM(p.amount) AS total_amount,
					c.name
				FROM personal_bot.t_payments p
				INNER JOIN personal_bot.t_currencies c
					ON p.currency_id = c.id
				WHERE p.description ~* $1
				AND p.user_id = $2
				AND p.last_updated >= $3::timestamptz
				AND p.last_updated <= $4::timestamptz
				GROUP BY DATE_TRUNC('month', p.last_updated), c.name
				ORDER BY DATE_TRUNC('month', p.last_updated);`

	rows, err := db.GetConnection().Query(statement, matcher, userId, from, to)

	if err != nil {
		logger.Error("Payments Repository - Get Summary Of Certain Payments By Dates And User Id", err.Error())

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment PaymentSummary = PaymentSummary{
			Currency: &Currency{},
		}

		if err := rows.Scan(&payment.Month, &payment.TotalCount, &payment.TotalAmount, &payment.Currency.Name); err != nil {
			return payments, err
		}

		payments.PaymentsSummary = append(payments.PaymentsSummary, &payment)
	}

	return payments, nil
}

// Return all the payments made by an user in a date range
func GetPaymentsByUserId(userId int) (*UserPayments, error) {
	var payments *UserPayments = &UserPayments{
		Payments: make([]*Payment, 0),
	}

	user, err := GetUser(userId)

	if err != nil {
		logger.Error("Currencies Repository - Get Payments - Get Users", err.Error())

		return nil, err
	}

	payments.User = user

	statement := `SELECT 
					pay.id payment_id, 
					pay.amount, 
					pay.last_updated AT TIME ZONE 'UTC-6',
					pay.dolar_price,
					pay.tags,
					curr.name currency,
					pay.gmail_id,
					pay.description,
					pay.id_bank,
					COALESCE(bank.name, 'No Data Source Registered') bank_name
				FROM personal_bot.t_payments pay
				INNER JOIN personal_bot.t_currencies curr
					ON pay.currency_id = curr.id
				LEFT JOIN personal_bot.t_banks bank
					ON pay.id_bank = bank.id
				WHERE pay.user_id = $1
				ORDER BY pay.last_updated DESC`

	rows, err := db.GetConnection().Query(statement, userId)

	if err != nil {
		logger.Error("Currencies Repository - Get Payments", err.Error())

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment Payment = Payment{
			Bank: &Bank{},
		}
		var tagsBody string = ""

		if err := rows.Scan(&payment.Id, &payment.Amount, &payment.LastUpdated, &payment.DolarPrice, &tagsBody, &payment.Currency, &payment.GmailId, &payment.Description, &payment.IdBank, &payment.Bank.Name); err != nil {
			return payments, err
		}

		if payment.IdBank != nil {
			payment.Bank.Id = *payment.IdBank
		}

		var tags []string = make([]string, 0)

		err = json.Unmarshal([]byte(tagsBody), &tags)

		if err != nil {
			continue
		}

		payment.Tags = tags

		payments.Payments = append(payments.Payments, &payment)
	}

	return payments, nil
}

// GetBankByName returns a bank with the given name, using a local cache to avoid unnecessary database calls.
func GetBankByName(name string) (*Bank, error) {
	// First, try to get the bank from the cache.
	if cached, ok := bankCache[name]; ok {
		if cached.expiration.After(time.Now()) {
			return cached.bank, nil
		} else {
			delete(bankCache, name)
		}
	}

	// If the bank was not found in the cache, query the database.
	statement := "SELECT id, name FROM personal_bot.t_banks WHERE name = $1"
	var b Bank
	err := db.GetConnection().QueryRow(statement, name).Scan(&b.Id, &b.Name)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("bank not found")
	}

	if err != nil {
		logger.Error("Error querying bank by name", err.Error())
		return nil, err
	}

	// Store the bank in the cache for future requests.
	cacheItem := bankCacheItem{bank: &b, expiration: time.Now().Add(24 * time.Hour)}

	bankCache[name] = cacheItem

	return &b, nil
}
