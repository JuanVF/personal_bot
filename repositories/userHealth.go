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
	"errors"
)

type HealthCondition struct {
	Id          int
	Name        string
	Description string
}

type UserHealth struct {
	Id                int
	UserId            int
	HealthConditionId int
	HealthCondition   HealthCondition
	DiagnosisDate     string
	Treatment         string
	DischargedDate    string
}

// Get a health condition by name
func GetHealthConditionByName(name string) (*HealthCondition, error) {
	statement := "SELECT id, name, description FROM personal_bot.t_health_conditions WHERE name = $1"

	row := db.GetConnection().QueryRow(statement, name)

	var healthCondition HealthCondition

	if err := row.Scan(&healthCondition.Id, &healthCondition.Name, &healthCondition.Description); err != nil {
		logger.Error("Health Condition Repository - Get Health Condition By Name", err.Error())

		return nil, err
	}

	return &healthCondition, nil
}

// Create a new user health record in the DB
func CreateUserHealth(userHealth *UserHealth) error {
	statement := `
	INSERT INTO personal_bot.t_user_health
	(user_id, health_condition_id, diagnosis_date, treatment, discharged_date)
	VALUES
	($1, $2, $3, $4, $5)`

	_, err := db.GetConnection().Exec(statement,
		userHealth.UserId,
		userHealth.HealthConditionId,
		userHealth.DiagnosisDate,
		userHealth.Treatment,
		userHealth.DischargedDate)

	if err != nil {
		logger.Error("User Health Repository - Create User Health", err.Error())
		return err
	}

	return nil
}

// Update an existing user health record in the DB
func UpdateUserHealth(userHealth *UserHealth) error {
	statement := `
	UPDATE personal_bot.t_user_health
	SET health_condition_id=$2, diagnosis_date=$3, treatment=$4, discharged_date=$5
	WHERE id=$1`

	result, err := db.GetConnection().Exec(statement,
		userHealth.Id,
		userHealth.HealthConditionId,
		userHealth.DiagnosisDate,
		userHealth.Treatment,
		userHealth.DischargedDate)

	if err != nil {
		logger.Error("User Health Repository - Update User Health", err.Error())
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}

// Get a user health record by user id
func GetUserHealthByUserId(userId int) ([]*UserHealth, error) {
	statement := `SELECT uh.id, uh.user_id, uh.health_condition_id, uh.diagnosis_date, uh.treatment, 
					  uh.discharged_date, 
					  hc.name as health_condition_name, hc.description as health_condition_description
				   FROM personal_bot.t_user_health uh 
				   INNER JOIN personal_bot.t_health_conditions hc ON uh.health_condition_id = hc.id
				   WHERE uh.user_id = $1`

	rows, err := db.GetConnection().Query(statement, userId)

	if err != nil {
		logger.Error("UserHealth Repository - GetUserHealthByUserId", err.Error())

		return nil, err
	}

	defer rows.Close()

	var userHealth []*UserHealth = make([]*UserHealth, 0)

	for rows.Next() {
		var uh UserHealth

		err := rows.Scan(&uh.Id, &uh.UserId, &uh.HealthConditionId, &uh.DiagnosisDate,
			&uh.Treatment, &uh.DischargedDate, &uh.HealthCondition.Name,
			&uh.HealthCondition.Description)

		if err != nil {
			logger.Error("UserHealth Repository - GetUserHealthByUserId", err.Error())

			return nil, err
		}

		userHealth = append(userHealth, &uh)
	}

	return userHealth, nil
}
