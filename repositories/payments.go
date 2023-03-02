package repositories

import "encoding/json"

type CreatePayment struct {
	Amount     float64
	UserId     int
	CurrencyId int
	Tags       []*Tag
}

type Payment struct {
	Id          int
	Amount      float64
	LastUpdated string
	Currency    string
	Tags        []*Tag
}

type UserPayments struct {
	User     *User
	Payments []*Payment
}

// Inserts a payment and its tags to the database
func InsertPayment(payment *CreatePayment) error {
	var paymentId int = 0

	statement := `INSERT INTO personal_bot.t_payments(
					amount, last_updated, user_id, currency_id)
				VALUES ($1, NOW(), $2, $3)
				RETURNING id`

	err := db.GetConnection().QueryRow(statement, payment.Amount, payment.UserId, payment.CurrencyId).Scan(&paymentId)

	if err != nil {
		logger.Error("Payment Repository - Insert Payment", err.Error())

		return err
	}

	return InsertTagsPerPayment(paymentId, payment.Tags)
}

// Return all the payments made by an user in a date range
func GetPaymentsByUserIdAndDateRange(userId int, start, end string) (*UserPayments, error) {
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
                    pay.last_updated,
                    curr.name currency,
					(SELECT 
						json_agg(t)
					FROM (
						SELECT 
							tags.id,
							tags.name
						FROM personal_bot.t_payments_per_tags ppt
						INNER JOIN personal_bot.t_tags tags
							ON ppt.id_tag = tags.id
						WHERE ppt.id_payment = pay.id
					) AS t) tags
                FROM personal_bot.t_payments pay
                INNER JOIN personal_bot.t_currencies curr
                    ON pay.currency_id = curr.id
                WHERE pay.user_id = $1 AND pay.last_updated BETWEEN $2 AND $3`

	rows, err := db.GetConnection().Query(statement, userId, start, end)

	if err != nil {
		logger.Error("Currencies Repository - Get Payments", err.Error())

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment Payment = Payment{}
		var tags string = ""

		if err := rows.Scan(&payment.Id, &payment.Amount, &payment.LastUpdated, &payment.Currency, &tags); err != nil {
			return payments, err
		}

		if err = json.Unmarshal([]byte(tags), &payment.Tags); err != nil {
			return payments, err
		}

		payments.Payments = append(payments.Payments, &payment)
	}

	return payments, nil
}
