package repositories

import "encoding/json"

type CreatePayment struct {
	Amount      float64
	UserId      int
	CurrencyId  int
	DolarPrice  float64
	Tags        []string
	LastUpdated string
}

type Payment struct {
	Id          int
	Amount      float64
	LastUpdated string
	Currency    string
	DolarPrice  float64
	Tags        []string
}

type UserPayments struct {
	User     *User
	Payments []*Payment
}

// Inserts a payment and its tags to the database
func InsertPayment(payment *CreatePayment) (int, error) {
	var paymentId int = 0

	statement := `INSERT INTO personal_bot.t_payments(
					amount, last_updated, user_id, currency_id, dolar_price, tags)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`

	tags, err := json.Marshal(payment.Tags)

	if err != nil {
		return 0, err
	}

	err = db.GetConnection().QueryRow(statement, payment.Amount, payment.LastUpdated, payment.UserId, payment.CurrencyId, payment.DolarPrice, string(tags)).Scan(&paymentId)

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
                    curr.name currency
                FROM personal_bot.t_payments pay
                INNER JOIN personal_bot.t_currencies curr
                    ON pay.currency_id = curr.id
                WHERE pay.user_id = $1`

	rows, err := db.GetConnection().Query(statement, userId)

	if err != nil {
		logger.Error("Currencies Repository - Get Payments", err.Error())

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment Payment = Payment{}
		var tagsBody string = ""

		if err := rows.Scan(&payment.Id, &payment.Amount, &payment.LastUpdated, &payment.DolarPrice, &tagsBody, &payment.Currency); err != nil {
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
