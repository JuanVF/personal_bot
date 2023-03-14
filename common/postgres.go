package common

import (
	"database/sql"
	"fmt"

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

	db, err := sql.Open("postgres", d.getDriver())

	if err != nil {
		GetLogger().Error("Postgres DB", err.Error())

		panic(err)
	}

	d.connection = db

	return db
}

// Returns the driver to connect to Postgres
func (d *DB) getDriver() string {
	pbConf := GetConfig().PersonalBotDB

	GetLogger().LogObject("Common", pbConf)

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", pbConf.Host, pbConf.Port, pbConf.User, pbConf.Password, pbConf.DBName, pbConf.SSLMode)
}
