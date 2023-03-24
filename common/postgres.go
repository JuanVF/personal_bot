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
		d.connection = d.getTestEnvConnection()
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

// getTestEnvConnection Use a mock database connection for tests
func (d *DB) getTestEnvConnection() *sql.DB {
	// Use a mock database connection for tests
	mockDB, _, err := sqlmock.New()

	if err != nil {
		GetLogger().Error("Mock DB", err.Error())
		panic(err)
	}

	return mockDB
}

// Returns the driver to connect to Postgres
func (d *DB) getDriver() string {
	pbConf := GetConfig().PersonalBotDB

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", pbConf.Host, pbConf.Port, pbConf.User, pbConf.Password, pbConf.DBName, pbConf.SSLMode)
}
