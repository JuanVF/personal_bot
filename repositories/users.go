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
	"fmt"
)

type User struct {
	Id              int
	Name            string
	LastName        string
	GoogleMe        string
	LastUpdated     string
	Weight          float64
	Height          float64
	ActivityLevelId int
	ActivityLevel   ActivityLevel
}

type CreateUserBody struct {
	Name            string
	LastName        string
	GoogleMe        string
	Weight          float64
	Height          float64
	ActivityLevelId int
}

type ActivityLevel struct {
	Id          int
	Name        string
	Description string
}

// Get an activity level by name
func GetActivityLevelByName(name string) (*ActivityLevel, error) {
	var activityLevel ActivityLevel

	query := "SELECT id, name, description FROM personal_bot.t_activity_levels WHERE name = $1"

	err := db.GetConnection().QueryRow(query, name).Scan(&activityLevel.Id, &activityLevel.Name, &activityLevel.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Not Found")
		} else {
			logger.Error("Activity Level Repository - Get Activity Level By Name", err.Error())

			return nil, err
		}
	}

	return &activityLevel, nil
}

// Returns a single user from DB
func GetUser(id int) (*User, error) {
	var user User = User{}

	statement := `SELECT 
						u.id, u.name, u.last_name, u.google_me, u.last_updated, u.weight, u.height, u.activity_level_id,
						al.name, al.description
					FROM personal_bot.t_users u
					INNER JOIN personal_bot.t_activity_levels al ON al.id = u.activity_level_id
					WHERE u.id = $1`

	err := db.GetConnection().
		QueryRow(statement, id).Scan(&user.Id, &user.Name, &user.LastName, &user.GoogleMe, &user.LastUpdated, &user.Weight, &user.Height, &user.ActivityLevelId, &user.ActivityLevel.Name, &user.ActivityLevel.Description)

	if err != nil {
		logger.Error("User Repository - Get User", err.Error())

		return nil, err
	}

	user.ActivityLevel.Id = user.ActivityLevelId

	return &user, nil
}

// Returns a single user from DB
func GetUserByGoogleMe(email string) (*User, error) {
	var user User = User{}

	statement := `SELECT 
					u.id, u.name, u.last_name, u.google_me, u.last_updated, u.weight, u.height, u.activity_level_id,
					al.name, al.description
				  FROM personal_bot.t_users u
				  INNER JOIN personal_bot.t_activity_levels al ON al.id = u.activity_level_id
				  WHERE u.google_me = $1`

	err := db.GetConnection().
		QueryRow(statement, email).
		Scan(&user.Id, &user.Name, &user.LastName, &user.GoogleMe, &user.LastUpdated, &user.Weight, &user.Height, &user.ActivityLevelId, &user.ActivityLevel.Name, &user.ActivityLevel.Description)

	if err != nil {
		logger.Error("User Repository - Get User By Google Me", err.Error())

		return nil, err
	}

	user.ActivityLevel.Id = user.ActivityLevelId

	return &user, nil
}

// Creates an user in DB
func CreateUser(user *CreateUserBody) (int, error) {
	statement := `INSERT INTO personal_bot.t_users(
					name, last_name, google_me, last_updated, weight, height, activity_level_id)
				VALUES 
					($1, $2, $3, NOW(), $4, $5, $6) 
				RETURNING id`

	var id int

	err := db.GetConnection().QueryRow(statement, user.Name, user.LastName, user.GoogleMe, user.Weight, user.Height, user.ActivityLevelId).Scan(&id)

	if err != nil {
		logger.Error("Bot Repository - CreateUser", err.Error())

		return 0, err
	}

	return id, nil
}
