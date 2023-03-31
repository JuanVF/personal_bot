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

import "encoding/json"

type CreatePayment struct {
	Amount      float64
	UserId      int
	CurrencyId  int
	DolarPrice  float64
	Tags        []string
	LastUpdated string
	GmailId     string
	Description string
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
}

type UserPayments struct {
	User     *User
	Payments []*Payment
}

// Inserts a payment and its tags to the database
func InsertPayment(payment *CreatePayment) (int, error) {
	var paymentId int = 0

	statement := `INSERT INTO personal_bot.t_payments(
					amount, last_updated, user_id, currency_id, dolar_price, tags, gmail_id, description)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id`

	tags, err := json.Marshal(payment.Tags)

	if err != nil {
		return 0, err
	}

	err = db.GetConnection().QueryRow(statement, payment.Amount, payment.LastUpdated, payment.UserId, payment.CurrencyId, payment.DolarPrice, string(tags), payment.GmailId, payment.Description).Scan(&paymentId)

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
					pay.description
                FROM personal_bot.t_payments pay
                INNER JOIN personal_bot.t_currencies curr
                    ON pay.currency_id = curr.id
                WHERE pay.user_id = $1
				ORDER BY pay.last_updated DESC`

	rows, err := db.GetConnection().Query(statement, userId)

	if err != nil {
		logger.Error("Currencies Repository - Get Payments", err.Error())

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment Payment = Payment{}
		var tagsBody string = ""

		if err := rows.Scan(&payment.Id, &payment.Amount, &payment.LastUpdated, &payment.DolarPrice, &tagsBody, &payment.Currency, &payment.GmailId, &payment.Description); err != nil {
			return payments, err
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
