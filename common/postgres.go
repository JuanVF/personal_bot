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
package common

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

var db *DB = nil

// Returns the instance for the DB
func GetDB() *DB {
	if db == nil {
		db = &DB{}
	}

	return db
}

// Returns the DB Connection
func (d *DB) GetConnection() *sql.DB {
	if d.connection != nil {
		return d.connection
	}

	env := GetEnvironment()

	if env == "test" {
		mockDB, mock := d.getTestEnvConnection()

		d.connection = mockDB
		d.mock = mock
	} else {
		d.connection = d.getConnection()
	}

	return d.connection
}

// getConnection Use a real database connection for non-test environments
func (d *DB) getConnection() *sql.DB {
	db, err := sql.Open("postgres", d.getDriver())

	if err != nil {
		GetLogger().Error("Postgres DB", err.Error())

		panic(err)
	}

	return db
}

// getMock Use a mock database connection for tests. USE ONLY FOR TESTS
func (d *DB) GetMock() *sqlmock.Sqlmock {
	return d.mock
}

// getTestEnvConnection Use a mock database connection for tests
func (d *DB) getTestEnvConnection() (*sql.DB, *sqlmock.Sqlmock) {
	// Use a mock database connection for tests
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		GetLogger().Error("Mock DB", err.Error())
		panic(err)
	}

	return mockDB, &mock
}

// Returns the driver to connect to Postgres
func (d *DB) getDriver() string {
	pbConf := GetConfig().PersonalBotDB

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", pbConf.Host, pbConf.Port, pbConf.User, pbConf.Password, pbConf.DBName, pbConf.SSLMode)
}
