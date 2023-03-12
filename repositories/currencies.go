package repositories

type Currency struct {
	Id   int
	Name string
}

var currenciesCache []*Currency = make([]*Currency, 0)

// Query all the currencies from the DB
func GetCurrencies() ([]*Currency, error) {
	if len(currenciesCache) != 0 {
		return currenciesCache, nil
	}

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

	currenciesCache = currencies

	return currencies, nil
}

// Return a currency by its name
func GetCurrencyByName(name string) *Currency {
	currencies, err := GetCurrencies()

	if err != nil {
		return nil
	}

	for _, c := range currencies {
		if c.Name == name {
			return c
		}
	}

	return nil
}
