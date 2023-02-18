package repositories

type Currency struct {
	Id   int
	Name string
}

// Query all the currencies from the DB
func GetCurrencies() ([]*Currency, error) {
	statement := "SELECT id, name FROM personal_bot.t_currencies"

	rows, err := db.GetConnection().Query(statement)

	if err != nil {
		logger.Error("Currencies Repository - Get Currencies", err.Error())

		return []*Currency{}, err
	}

	defer rows.Close()

	var currencies []*Currency = make([]*Currency, 0)

	for rows.Next() {
		var currency Currency

		if err := rows.Scan(&currency.Id, &currency.Name); err != nil {
			return currencies, err
		}

		currencies = append(currencies, &currency)
	}

	return currencies, nil
}
