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
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var userTest *CreateUserBody = &CreateUserBody{
	Name:            "Jhon",
	LastName:        "Doe",
	GoogleMe:        "jhon.doe@gmail.com",
	Weight:          80.5,
	Height:          180,
	ActivityLevelId: 1,
}

var fakeUserTest *CreateUserBody = &CreateUserBody{
	Name:            "Not Jhon",
	LastName:        "Fake Doe",
	GoogleMe:        "fake.jhon.doe@gmail.com",
	Weight:          62.1,
	Height:          172,
	ActivityLevelId: 2,
}

func TestCreateUser(t *testing.T) {
	tName := "User Repository - Create User"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(`INSERT INTO personal_bot.t_users(
						name, last_name, google_me, last_updated, weight, height, activity_level_id)
					VALUES 
						($1, $2, $3, NOW(), $4, $5, $6) 
					RETURNING id`).
		WithArgs(userTest.Name, userTest.LastName, userTest.GoogleMe, userTest.Weight, userTest.Height, userTest.ActivityLevelId).
		WillReturnRows(rows)

	// Call function to be tested
	id, _ := CreateUser(userTest)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if id != 1 {
		logger.TestError(tName, "User created with id: 1", fmt.Sprintf("User created with id: %d", id), t)
	}
}

func TestGetUserByGoogleMe(t *testing.T) {
	tName := "User Repository - Get User By Google Me"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated", "weight", "height", "activity_level_id", "name", "description"}).
		AddRow(1, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00", userTest.Weight, userTest.Height, userTest.ActivityLevelId, "Sedentary", "Little or no exercise")

	mock.ExpectQuery(`SELECT 
						u.id, u.name, u.last_name, u.google_me, u.last_updated, u.weight, u.height, u.activity_level_id,
						al.name, al.description
					FROM personal_bot.t_users u
					INNER JOIN personal_bot.t_activity_levels al ON al.id = u.activity_level_id
					WHERE u.google_me = $1`).
		WithArgs(userTest.GoogleMe).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUserByGoogleMe(userTest.GoogleMe)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user.GoogleMe != userTest.GoogleMe {
		logger.TestError(tName, fmt.Sprintf("User with google_me: %s", userTest.GoogleMe), fmt.Sprintf("User with google_me: %s", user.GoogleMe), t)
	}

	if user.Name != userTest.Name {
		logger.TestError(tName, fmt.Sprintf("User with name: %s", userTest.Name), fmt.Sprintf("User with name: %s", user.Name), t)
	}

	if user.LastName != userTest.LastName {
		logger.TestError(tName, fmt.Sprintf("User with last name: %s", userTest.LastName), fmt.Sprintf("User with last name: %s", user.LastName), t)
	}
}

func TestGetUserByGoogleMeWhenItShouldNotExists(t *testing.T) {
	tName := "User Repository - Get User By Google Me - When it should not exists"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated", "weight", "height", "activity_level_id", "name", "description"})

	mock.ExpectQuery(`SELECT 
						u.id, u.name, u.last_name, u.google_me, u.last_updated, u.weight, u.height, u.activity_level_id,
						al.name, al.description
					FROM personal_bot.t_users u
					INNER JOIN personal_bot.t_activity_levels al ON al.id = u.activity_level_id
					WHERE u.google_me = $1`).
		WithArgs(fakeUserTest.GoogleMe).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUserByGoogleMe(fakeUserTest.GoogleMe)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user != nil {
		logger.TestError(tName, "User should be nil", fmt.Sprintf("User with google_me: %s", user.GoogleMe), t)
	}
}

func TestGetUser(t *testing.T) {
	tName := "User Repository - Get User"

	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated", "weight", "height", "activity_level_id", "name", "description"}).
		AddRow(1, userTest.Name, userTest.LastName, userTest.GoogleMe, "2019-01-01 00:00:00", userTest.Weight, userTest.Height, userTest.ActivityLevelId, "Sedentary", "Little or no exercise")

	mock.ExpectQuery(`SELECT 
						u.id, u.name, u.last_name, u.google_me, u.last_updated, u.weight, u.height, u.activity_level_id,
						al.name, al.description
					FROM personal_bot.t_users u
					INNER JOIN personal_bot.t_activity_levels al ON al.id = u.activity_level_id
					WHERE u.id = $1`).
		WithArgs(1).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUser(1)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user == nil {
		logger.TestError(tName, "User should not be nil", "User is nil", t)
	}

	if user.Id != 1 {
		logger.TestError(tName, fmt.Sprintf("User id should be 1, got %d", user.Id), "", t)
	}

	if user.Name != userTest.Name {
		logger.TestError(tName, fmt.Sprintf("User with name: %s", userTest.Name), fmt.Sprintf("User with name: %s", user.Name), t)
	}

	if user.LastName != userTest.LastName {
		logger.TestError(tName, fmt.Sprintf("User with last name: %s", userTest.LastName), fmt.Sprintf("User with last name: %s", user.LastName), t)
	}

	if user.GoogleMe != userTest.GoogleMe {
		logger.TestError(tName, fmt.Sprintf("User with google_me: %s", userTest.GoogleMe), fmt.Sprintf("User with google_me: %s", user.GoogleMe), t)
	}
}

func TestGetUserWhenNotFound(t *testing.T) {
	tName := "User Repository - Get User - When not found"
	mock := *db.GetMock()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "google_me", "last_updated", "weight", "height", "activity_level_id", "name", "description"})

	mock.ExpectQuery(`SELECT 
						u.id, u.name, u.last_name, u.google_me, u.last_updated, u.weight, u.height, u.activity_level_id,
						al.name, al.description
					FROM personal_bot.t_users u
					INNER JOIN personal_bot.t_activity_levels al ON al.id = u.activity_level_id
					WHERE u.id = $1`).
		WithArgs(100).
		WillReturnRows(rows)

	// Call function to be tested
	user, _ := GetUser(100)

	// Verify that the expected query was executed
	err := mock.ExpectationsWereMet()

	if err != nil {
		logger.TestError(tName, "Expectations Fullfilled", err.Error(), t)
	}

	if user != nil {
		logger.TestError(tName, "User should be nil", fmt.Sprintf("User: %v", user), t)
	}
}
