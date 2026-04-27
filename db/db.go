package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Could Not Connect to DB: " + err.Error())
	}

	createTables()

	DB.SetMaxOpenConns(10)

	DB.SetMaxIdleConns(5)

}

func createTables() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could Not Create Users Table")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventTable)

	if err != nil {
		println(err)
		panic("Could Not Create Events Table")
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id)
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		println(err)
		panic("Could Not Create Registrations Table")
	}
}
