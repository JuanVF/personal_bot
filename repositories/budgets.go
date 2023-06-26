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

type Budget struct {
	Id          int
	Name        string
	UserId      int
	User        *User
	Amount      float64
	Matcher     string
	LastUpdated string
}

// Query the budgets from an user by id
func GetBudgetsByUserId(userId int) ([]*Budget, error) {
	statement := "SELECT id, name, amount, matcher, last_updated FROM personal_bot.t_budgets WHERE user_id = $1"

	rows, err := db.GetConnection().Query(statement, userId)

	if err != nil {
		logger.Error("Budget Repository - Get Budget By User Id", err.Error())

		return []*Budget{}, err
	}

	defer rows.Close()

	var budgets []*Budget = make([]*Budget, 0)

	for rows.Next() {
		var budget Budget

		if err := rows.Scan(&budget.Id, &budget.Name, &budget.Amount, &budget.Matcher, &budget.LastUpdated); err != nil {
			return budgets, err
		}

		budgets = append(budgets, &budget)
	}

	return budgets, nil
}
