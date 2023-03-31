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
