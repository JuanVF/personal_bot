package repositories

type Balance struct {
	Id          int
	UserId      int
	Amount      float64
	Expenses    float64
	Currency    *Currency
	ExpiresIn   string
	LastUpdated string
}

// Retrieves the Token Info for a certain user
func GetBalanceByUserId(userId int) (*Balance, error) {
	var balance *Balance = &Balance{
		UserId:   userId,
		Currency: &Currency{},
	}

	statement := `SELECT 
                    bal.id, 
                    bal.amount, 
                    bal.expenses, 
                    bal.last_updated,
                    bal.currency_id,
                    cur.name currency_name
                FROM personal_bot.t_balances bal
                INNER JOIN personal_bot.t_currencies cur
                ON cur.id = bal.currency_id
                WHERE bal.user_id = $1;`

	err := db.GetConnection().QueryRow(statement, userId).Scan(
		&balance.Id,
		&balance.Amount,
		&balance.Expenses,
		&balance.LastUpdated,
		&balance.Currency.Id,
		&balance.Currency.Name,
	)

	if err != nil {
		logger.Error("Balance Repository - Get Balance By User Id", err.Error())

		return nil, err
	}

	return balance, nil
}

// Updates the balance
func UpdateBalance(balance *Balance) error {
	statement := `UPDATE personal_bot.t_balances
                SET amount=$1, expenses=$2, last_updated=NOW()
                WHERE user_id = $3;`

	_, err := db.GetConnection().Exec(statement, balance.Amount, balance.Expenses, balance.UserId)

	if err != nil {
		logger.Error("Balance Repository - Update Token", err.Error())
	}

	return err
}
